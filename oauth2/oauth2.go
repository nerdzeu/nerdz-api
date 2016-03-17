package oauth2

import (
	"bytes"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"net/http"
	"net/url"
)

var oauth *osin.Server

// Init initialize OAuth2 Authorization server. This is the first function to call before using OAuth2
func Init(server *osin.Server) {
	oauth = server
}

// Authorize is the action of GET /authorize and POST /authorize when authentication is required
func Authorize() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := oauth.NewResponse()
		defer resp.Close()
		if ar := oauth.HandleAuthorizeRequest(resp, c.Request().(*standard.Request).Request); ar != nil {
			if user, err := nerdz.HandleLoginPage(ar, c); err != nil {
				return nil // HandleLoginPage handles errors as well
			} else {
				ar.UserData = user.Counter
				ar.Authorized = true
				oauth.FinishAuthorizeRequest(resp, c.Request().(*standard.Request).Request, ar)
			}
		}

		if resp.IsError && resp.InternalError != nil {
			return c.JSON(http.StatusInternalServerError, &rest.Response{
				HumanMessage: "Internal Server error",
				Message:      resp.InternalError.Error(),
				Status:       http.StatusInternalServerError,
				Success:      false,
			})
		}

		return osin.OutputJSON(resp, c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
	}
}

// Token is the action of Get /token
func Token() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := oauth.NewResponse()
		defer resp.Close()

		if ar := oauth.HandleAccessRequest(resp, c.Request().(*standard.Request).Request); ar != nil {
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
			}
			oauth.FinishAccessRequest(resp, c.Request().(*standard.Request).Request, ar)
		}

		if resp.IsError && resp.InternalError != nil {
			return c.JSON(http.StatusInternalServerError, &rest.Response{
				HumanMessage: "Internal Server error",
				Message:      resp.InternalError.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		return osin.OutputJSON(resp, c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
	}
}

// Info is the action of GET /info
func Info() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := oauth.NewResponse()
		defer resp.Close()

		if ir := oauth.HandleInfoRequest(resp, c.Request().(*standard.Request).Request); ir != nil {
			oauth.FinishInfoRequest(resp, c.Request().(*standard.Request).Request, ir)
		}

		return osin.OutputJSON(resp, c.Response().(*standard.Response).ResponseWriter, c.Request().(*standard.Request).Request)
	}
}

// App is the application home endpoint (action of GET /app)
func App() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer
		buffer.WriteString("<html><body>")
		buffer.WriteString(fmt.Sprintf("<a href=\"authorize?response_type=code&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Code</a><br/>", url.QueryEscape(nerdz.Configuration.ApiURL().String()+"/appauth/code")))
		buffer.WriteString(fmt.Sprintf("<a href=\"authorize?response_type=token&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Implict</a><br/>", url.QueryEscape(nerdz.Configuration.ApiURL().String()+"/appauth/token")))
		buffer.WriteString(fmt.Sprintf("<a href=\"appauth/password\">Password</a><br/>"))
		buffer.WriteString(fmt.Sprintf("<a href=\"appauth/client_credentials\">Client Credentials</a><br/>"))
		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}

// Application destination - CODE. GET /appauth/code
func AppAuthCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request().(*standard.Request).Request

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
			url.QueryEscape(nerdz.Configuration.ApiURL().String()+"/appauth/code"), url.QueryEscape(code))

		// if parse, download and parse json
		if r.Form.Get("doparse") == "1" {
			err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.ApiURL().String()+"%s", aurl),
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
}

// AppAuthToken hanlde Application destination - TOKEN. GET /appauth/token
func AppAuthToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer

		buffer.WriteString("<html><body>")
		buffer.WriteString("APP AUTH - TOKEN<br/>")
		buffer.WriteString("Response data in fragment - not acessible via server - Nothing to do")
		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}

// AppAuthPassword handles Application destination - PASSWORD. GET /appauth/password
func AppAuthPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer
		buffer.WriteString("<html><body>")
		buffer.WriteString("APP AUTH - PASSWORD<br/>")

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("/token?grant_type=password&scope=everything&username=%s&password=%s",
			"test", "test")

		// download token
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.ApiURL().String()+"%s", aurl),
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
}

// AppAuthClientCredentials handles Application destination - CLIENT_CREDENTIALS. GET /appauth/client_credentials
func AppAuthClientCredentials() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer

		buffer.WriteString("<html><body>")
		buffer.WriteString("APP AUTH - CLIENT CREDENTIALS<br/>")

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("/token?grant_type=client_credentials")

		// download token
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.ApiURL().String()+"%s", aurl),
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
}

// AppAuthRefresh handles Application destination - REFRESH. GET /appauth/refresh
func AppAuthRefresh() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request().(*standard.Request).Request

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
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.ApiURL().String()+"%s", aurl),
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
}

// AppAuthInfo handles Application destination - INFO. GET /appauth/info
func AppAuthInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request().(*standard.Request).Request

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
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.ApiURL().String()+"%s", aurl),
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
}
