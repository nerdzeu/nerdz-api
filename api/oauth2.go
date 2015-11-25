package api

// Open url in browser:
// TODO: replace localhost:port with nerdz.Configuration.ApiHost (or whatever) and port everywhere
// http://localhost:14000/app

import (
	"bytes"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"net/http"
	"net/url"
)

// OAuth is the Authorization server, initialized on server start
var OAuth *osin.Server

// OAuth2Authorize is the action of GET /authorize and POST /authorize when authentication is required
func OAuth2Authorize(c *echo.Context) error {
	resp := OAuth.NewResponse()
	defer resp.Close()
	if ar := OAuth.HandleAuthorizeRequest(resp, c.Request()); ar != nil {
		if user, err := nerdz.HandleLoginPage(ar, c); err != nil {
			return nil // HandleLoginPage handles errors as well
		} else {
			ar.UserData = user.Counter
			ar.Authorized = true
			OAuth.FinishAuthorizeRequest(resp, c.Request(), ar)
		}
	}

	if resp.IsError && resp.InternalError != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			HumanMessage: "Internal Server error",
			Message:      resp.InternalError.Error(),
			Status:       http.StatusInternalServerError,
			Success:      false,
		})
	}

	return osin.OutputJSON(resp, c.Response().Writer(), c.Request())
}

// OAuth2Token is the action of Get /token
func OAuth2Token(c *echo.Context) error {
	resp := OAuth.NewResponse()
	defer resp.Close()

	if ar := OAuth.HandleAccessRequest(resp, c.Request()); ar != nil {
		switch ar.Type {
		case osin.AUTHORIZATION_CODE:
			ar.Authorized = true
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			if _, err := nerdz.Login(ar.Username, ar.Password); err == nil {
				ar.Authorized = true
			}
		case osin.CLIENT_CREDENTIALS:
			ar.Authorized = true
			/* TODO: support OAuth2 assertion flow?
			case osin.ASSERTION:
				if ar.AssertionType == "urn:osin.nerdz.complete" && ar.Assertion == "osin.data" {
					ar.Authorized = true
				}
			*/
		}
		OAuth.FinishAccessRequest(resp, c.Request(), ar)
	}

	if resp.IsError && resp.InternalError != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			HumanMessage: "Internal Server error",
			Message:      resp.InternalError.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	return osin.OutputJSON(resp, c.Response().Writer(), c.Request())
}

// OAuth2Info is the action of GET /info
func OAuth2Info(c *echo.Context) error {
	resp := OAuth.NewResponse()
	defer resp.Close()

	if ir := OAuth.HandleInfoRequest(resp, c.Request()); ir != nil {
		OAuth.FinishInfoRequest(resp, c.Request(), ir)
	}

	return osin.OutputJSON(resp, c.Response().Writer(), c.Request())
}

// OAuth2App is the application home endpoint (action of GET /app)
func OAuth2App(c *echo.Context) error {
	var buffer bytes.Buffer
	buffer.WriteString("<html><body>")
	buffer.WriteString(fmt.Sprintf("<a href=\"/authorize?response_type=code&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Code</a><br/>", url.QueryEscape("https://localhost:14000/appauth/code")))
	buffer.WriteString(fmt.Sprintf("<a href=\"/authorize?response_type=token&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Implict</a><br/>", url.QueryEscape("http://localhost:14000/appauth/token")))
	buffer.WriteString(fmt.Sprintf("<a href=\"/appauth/password\">Password</a><br/>"))
	buffer.WriteString(fmt.Sprintf("<a href=\"/appauth/client_credentials\">Client Credentials</a><br/>"))
	// TODO: assertion support?
	//buffer.WriteString(fmt.Sprintf("<a href=\"/appauth/assertion\">Assertion</a><br/>")))
	buffer.WriteString("</body></html>")
	return c.HTML(http.StatusOK, buffer.String())
}

// Application destination - CODE. GET /appauth/code
func OAuth2AppAuthCode(c *echo.Context) error {
	r := c.Request()
	r.ParseForm()
	code := r.Form.Get("code")

	var buffer bytes.Buffer
	buffer.WriteString("<html><body>")
	buffer.WriteString("APP AUTH - CODE<br/>")

	if code == "" {
		buffer.WriteString("Nothing to do</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=authorization_code&client_id=1234&client_secret=aabbccdd&state=xyz&redirect_uri=%s&code=%s",
		url.QueryEscape("http://localhost:14000/appauth/code"), url.QueryEscape(code))

	// if parse, download and parse json
	if r.Form.Get("doparse") == "1" {
		err := nerdz.DownloadAccessToken(fmt.Sprintf("http://localhost:14000%s", aurl),
			&osin.BasicAuth{"1234", "aabbccdd"}, jr)
		if err != nil {
			buffer.WriteString(err.Error())
			buffer.WriteString("<br/>")
		}
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		buffer.WriteString(fmt.Sprintf("ERROR: %s<br/>\n", erd))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		buffer.WriteString(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at))
	}

	buffer.WriteString(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr))

	// output links
	buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Goto Token URL</a><br/>", aurl))

	cururl := *r.URL
	curq := cururl.Query()
	curq.Add("doparse", "1")
	cururl.RawQuery = curq.Encode()
	buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Download Token</a><br/>", cururl.String()))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
	}

	buffer.WriteString("</body></html>")

	return c.HTML(http.StatusOK, buffer.String())
}

// OAuth2AppAuthToken hanlde Application destination - TOKEN. GET /appauth/token
func OAuth2AppAuthToken(c *echo.Context) error {
	var buffer bytes.Buffer
	r := c.Request()
	r.ParseForm()
	buffer.WriteString("<html><body>")
	buffer.WriteString("APP AUTH - TOKEN<br/>")
	buffer.WriteString("Response data in fragment - not acessible via server - Nothing to do")
	buffer.WriteString("</body></html>")
	return c.HTML(http.StatusOK, buffer.String())
}

// OAuth2AppAuthPassword handles Application destination - PASSWORD. GET /appauth/password
func OAuth2AppAuthPassword(c *echo.Context) error {
	r := c.Request()
	r.ParseForm()

	var buffer bytes.Buffer
	buffer.WriteString("<html><body>")
	buffer.WriteString("APP AUTH - PASSWORD<br/>")

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=password&scope=everything&username=%s&password=%s",
		"test", "test")

	// download token
	err := nerdz.DownloadAccessToken(fmt.Sprintf("http://localhost:14000%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		buffer.WriteString(err.Error())
		buffer.WriteString("<br/>")
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		buffer.WriteString(fmt.Sprintf("ERROR: %s<br/>\n", erd))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		buffer.WriteString(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at))
	}

	buffer.WriteString(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
	}

	buffer.WriteString("</body></html>")
	return c.HTML(http.StatusOK, buffer.String())
}

// OAuth2AppAuthClientCredentials handles Application destination - CLIENT_CREDENTIALS. GET /appauth/client_credentials
func OAuth2AppAuthClientCredentials(c *echo.Context) error {
	r := c.Request()
	r.ParseForm()

	var buffer bytes.Buffer

	buffer.WriteString("<html><body>")
	buffer.WriteString("APP AUTH - CLIENT CREDENTIALS<br/>")

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=client_credentials")

	// download token
	err := nerdz.DownloadAccessToken(fmt.Sprintf("http://localhost:14000%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		buffer.WriteString(err.Error())
		buffer.WriteString("<br/>")
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		buffer.WriteString(fmt.Sprintf("ERROR: %s<br/>\n", erd))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		buffer.WriteString(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at))
	}

	buffer.WriteString(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
	}

	buffer.WriteString("</body></html>")
	return c.HTML(http.StatusOK, buffer.String())
}

// OAuth2AppAuthRefresh handles Application destination - REFRESH. GET /appauth/refresh
func OAuth2AppAuthRefresh(c *echo.Context) error {
	r := c.Request()
	r.ParseForm()

	var buffer bytes.Buffer
	buffer.WriteString("<html><body>")
	buffer.WriteString("APP AUTH - REFRESH<br/>")

	code := r.Form.Get("code")

	if code == "" {
		buffer.WriteString("Nothing to do</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=refresh_token&refresh_token=%s", url.QueryEscape(code))

	// download token
	err := nerdz.DownloadAccessToken(fmt.Sprintf("http://localhost:14000%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		buffer.WriteString(err.Error())
		buffer.WriteString("<br/>")
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		buffer.WriteString(fmt.Sprintf("ERROR: %s<br/>\n", erd))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		buffer.WriteString(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at))
	}

	buffer.WriteString(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
	}

	buffer.WriteString("</body></html>")
	return c.HTML(http.StatusOK, buffer.String())
}

// OAuth2AppAuthInfo handles Application destination - INFO. GET /appauth/info
func OAuth2AppAuthInfo(c *echo.Context) error {
	r := c.Request()
	r.ParseForm()

	var buffer bytes.Buffer
	buffer.WriteString("<html><body>")
	buffer.WriteString("APP AUTH - INFO<br/>")

	code := r.Form.Get("code")

	if code == "" {
		buffer.WriteString("Nothing to do</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/info?code=%s", url.QueryEscape(code))

	// download token
	err := nerdz.DownloadAccessToken(fmt.Sprintf("http://localhost:14000%s", aurl),
		&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		buffer.WriteString(err.Error())
		buffer.WriteString("<br/>")
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		buffer.WriteString(fmt.Sprintf("ERROR: %s<br/>\n", erd))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		buffer.WriteString(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at))
	}

	buffer.WriteString(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
	}

	buffer.WriteString("</body></html>")
	return c.HTML(http.StatusOK, buffer.String())
}

/* TODO: OAuth2 assertion flow support?
// Application destination - ASSERTION
http.HandleFunc("/appauth/assertion", func(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	buffer.WriteString("<html><body>"))
	buffer.WriteString("APP AUTH - ASSERTION<br/>"))

	jr := make(map[string]interface{})

	// build access code url
	aurl := fmt.Sprintf("/token?grant_type=assertion&assertion_type=urn:osin.nerdz.complete&assertion=osin.data")

	// download token
	err := nerdz.DownloadAccessToken(fmt.Sprintf("http://localhost:14000%s", aurl),
	&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
	if err != nil {
		buffer.WriteString(err.Error()))
		buffer.WriteString("<br/>"))
	}

	// show json error
	if erd, ok := jr["error"]; ok {
		buffer.WriteString(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
	}

	// show json access token
	if at, ok := jr["access_token"]; ok {
		buffer.WriteString(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
	}

	buffer.WriteString(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

	if rt, ok := jr["refresh_token"]; ok {
		rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
	}

	if at, ok := jr["access_token"]; ok {
		rurl := fmt.Sprintf("/appauth/info?code=%s", at)
		buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
	}

	buffer.WriteString("</body></html>"))
}
*/
