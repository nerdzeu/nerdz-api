/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package appauth

import (
	"bytes"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"net/http"
	"net/url"
)

// Code is the Application destination - CODE. GET /oauth2/appauth/code
func Code() echo.HandlerFunc {
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
			url.QueryEscape(nerdz.Configuration.APIURL().String()+"/oauth2/appauth/code"), url.QueryEscape(code))

		// if parse, download and parse json
		if r.Form.Get("doparse") == "1" {
			err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.APIURL().String()+"%s", aurl),
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
			rurl := fmt.Sprintf("/oauth2/appauth/refresh?code=%s", rt)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/oauth2/appauth/info?code=%s", at)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
		}

		buffer.WriteString("</body></html>")

		return c.HTML(http.StatusOK, buffer.String())
	}
}

// Token hanlde Application destination - TOKEN. GET /oauth2/appauth/token
func Token() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer

		buffer.WriteString("<html><body>")
		buffer.WriteString("APP AUTH - TOKEN<br/>")
		buffer.WriteString("Response data in fragment - not acessible via server - Nothing to do")
		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}

// Password handles Application destination - PASSWORD. GET /oauth2/appauth/password
func Password() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer
		buffer.WriteString("<html><body>")
		buffer.WriteString("APP AUTH - PASSWORD<br/>")

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("/token?grant_type=password&scope=everything&username=%s&password=%s",
			"test", "test")

		// download token
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.APIURL().String()+"%s", aurl),
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
			rurl := fmt.Sprintf("/oauth2/appauth/refresh?code=%s", rt)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/oauth2/appauth/info?code=%s", at)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
		}

		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}

// ClientCredentials handles Application destination - CLIENT_CREDENTIALS. GET /oauth2/appauth/client_credentials
func ClientCredentials() echo.HandlerFunc {
	return func(c echo.Context) error {
		var buffer bytes.Buffer

		buffer.WriteString("<html><body>")
		buffer.WriteString("APP AUTH - CLIENT CREDENTIALS<br/>")

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("/token?grant_type=client_credentials")

		// download token
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.APIURL().String()+"%s", aurl),
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
			rurl := fmt.Sprintf("/oauth2/appauth/refresh?code=%s", rt)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/oauth2/appauth/info?code=%s", at)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
		}

		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}

// Refresh handles Application destination - REFRESH. GET /oauth2/appauth/refresh
func Refresh() echo.HandlerFunc {
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
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.APIURL().String()+"%s", aurl),
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
			rurl := fmt.Sprintf("/oauth2/appauth/refresh?code=%s", rt)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/oauth2/appauth/info?code=%s", at)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl))
		}

		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}

// Info handles Application destination - INFO. GET /oauth2/appauth/info
func Info() echo.HandlerFunc {
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
		err := nerdz.DownloadAccessToken(fmt.Sprintf(nerdz.Configuration.APIURL().String()+"%s", aurl),
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
			rurl := fmt.Sprintf("/oauth2/appauth/refresh?code=%s", rt)
			buffer.WriteString(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl))
		}

		buffer.WriteString("</body></html>")
		return c.HTML(http.StatusOK, buffer.String())
	}
}
