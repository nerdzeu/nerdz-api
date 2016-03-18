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

// Authorize is the action of GET /oauth2/authorize and POST /oauth2/authorize when authentication is required
func Authorize() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := oauth.NewResponse()
		defer resp.Close()

		if ar := oauth.HandleAuthorizeRequest(resp, c.Request().(*standard.Request).Request); ar != nil {
			user, err := nerdz.HandleLoginPage(ar, c)
			if err != nil {
				return nil // HandleLoginPage handles errors as well
			}
			ar.UserData = user.Counter
			ar.Authorized = true
			oauth.FinishAuthorizeRequest(resp, c.Request().(*standard.Request).Request, ar)
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

// Token is the action of Get /oauth2/token
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

// Info is the action of GET /oauth2/info
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

// App is the application home endpoint (action of GET /oauth2/app)
func App() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer
		buffer.WriteString("<html><body>")
		buffer.WriteString(fmt.Sprintf("<a href=\"authorize?response_type=code&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Code</a><br/>", url.QueryEscape(nerdz.Configuration.ApiURL().String()+"/oauth2/appauth/code")))
		buffer.WriteString(fmt.Sprintf("<a href=\"authorize?response_type=token&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Implict</a><br/>", url.QueryEscape(nerdz.Configuration.ApiURL().String()+"/oauth2/appauth/token")))
		buffer.WriteString(fmt.Sprintf("<a href=\"appauth/password\">Password</a><br/>"))
		buffer.WriteString(fmt.Sprintf("<a href=\"appauth/client_credentials\">Client Credentials</a><br/>"))
		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}
