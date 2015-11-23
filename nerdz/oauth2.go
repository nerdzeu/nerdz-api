package nerdz

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/utils"
)

// isValidScope checks if scope is a valid scope
func (s *OAuth2Storage) isValidScope(scope string) error {
	scopes := strings.Split(scope, " ")
	for _, s := range scopes {
		if !utils.InSlice(s, Configuration.Scopes) {
			return errors.New("Requested scope (" + s + ") does not exist")
		}
	}
	return nil
}

// OAuth2Storage implements osin.Storage interface
type OAuth2Storage struct {
}

// Clone the storage if needed. For example, using mgo, you can clone the session with session.Clone
// to avoid concurrent access problems.
// This is to avoid cloning the connection at each method access.
// Can return itself if not a problem.
func (s *OAuth2Storage) Clone() osin.Storage {
	return s
}

// Close the resources the Storage potentially holds (using Clone for example)
func (s *OAuth2Storage) Close() {
}

// GetClient loads the client by id (client_id)
func (s *OAuth2Storage) GetClient(id string) (osin.Client, error) {
	var err error
	var numericID uint64
	if numericID, err = strconv.ParseUint(id, 10, 64); err != nil {
		return nil, fmt.Errorf("invalid client_id: %s", id)
	}

	client := new(OAuth2Client)
	Db().First(client, numericID)
	if client.GetId() != id {
		return nil, errors.New("Client not found: " + id)
	}
	return client, nil
}

// SaveAuthorize saves authorize data.
func (s *OAuth2Storage) SaveAuthorize(data *osin.AuthorizeData) error {
	var clientID uint64
	var err error
	if clientID, err = strconv.ParseUint(data.Client.GetId(), 10, 64); err != nil {
		return err
	}

	if err = s.isValidScope(data.Scope); err != nil {
		return err
	}

	d := &OAuth2AuthorizeData{
		ClientID: clientID,
		Code:     data.Code,
		// CreatedAt field is automatically filled by the dbms
		ExpiresIn:   uint64(data.ExpiresIn),
		RedirectURI: data.RedirectUri,
		Scope:       data.Scope,
		State:       data.State,
		UserID:      data.UserData.(uint64)}

	return Db().Create(d).Error
}

// LoadAuthorize looks up osin.AuthorizeData by a code.
// osin.Client information MUST be loaded together.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	authorize := new(OAuth2AuthorizeData)
	Db().Model(OAuth2AuthorizeData{}).Where(&OAuth2AuthorizeData{Code: code}).Find(authorize)
	if code != authorize.Code {
		return nil, errors.New("Authorization data not found")
	}

	authData := &osin.AuthorizeData{
		Code:        code,
		CreatedAt:   authorize.CreatedAt.Round(time.Second),
		ExpiresIn:   int32(authorize.ExpiresIn),
		RedirectUri: authorize.RedirectURI,
		Scope:       authorize.Scope,
		State:       authorize.State,
		UserData:    authorize.UserID}

	if authData.IsExpired() {
		return nil, errors.New("Authorization data expired")
	}

	if client, err := s.GetClient(strconv.FormatUint(authorize.ClientID, 10)); err == nil {
		authData.Client = client
		return authData, nil
	}
	return nil, errors.New("LoadAuthorize: Client not found")
}

// RemoveAuthorize revokes or deletes the authorization code.
func (s *OAuth2Storage) RemoveAuthorize(code string) error {
	return Db().Where(&OAuth2AuthorizeData{Code: code}).Delete(OAuth2AuthorizeData{}).Error
}

// SaveAccess writes osin.AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (s *OAuth2Storage) SaveAccess(accessData *osin.AccessData) error {
	var clientID uint64
	var err error
	if clientID, err = strconv.ParseUint(accessData.Client.GetId(), 10, 64); err != nil {
		return errors.New("Invalid Client ID")
	}

	if err = s.isValidScope(accessData.Scope); err != nil {
		return err
	}

	var accessDataIDPtr sql.NullInt64
	if accessData.AccessData != nil {
		var father OAuth2AccessData
		Db().Model(OAuth2AccessData{}).Where(&OAuth2AccessData{AccessToken: accessData.AccessData.AccessToken}).Find(&father)

		if father.ID == 0 {
			return errors.New("Error fetching parent Access Data ID")
		}

		accessDataIDPtr.Int64 = int64(father.ID)
		accessDataIDPtr.Valid = true
	}

	// required to fill the foreign key
	var authorizeData OAuth2AuthorizeData
	Db().Model(OAuth2AuthorizeData{}).Where(&OAuth2AuthorizeData{Code: accessData.AuthorizeData.Code}).Find(&authorizeData)
	if authorizeData.ID == 0 {
		return fmt.Errorf("SaveAccess: can't load authorize data with code: %s", accessData.AuthorizeData.Code)
	}

	tx := Db().Begin()

	var refreshTokenFK sql.NullInt64

	oauthAccessData := &OAuth2AccessData{
		AccessDataID:    accessDataIDPtr,
		AccessToken:     accessData.AccessToken,
		AuthorizeDataID: authorizeData.ID,
		ClientID:        clientID,
		//CreatedAt:       accessData.CreatedAt, <- dbms handled
		ExpiresIn:   uint64(accessData.ExpiresIn),
		RedirectURI: accessData.RedirectUri,
		Scope:       accessData.Scope,
		UserID:      accessData.UserData.(uint64)}

	if accessData.RefreshToken != "" {
		// Create refresh token
		var newRefreshToken OAuth2RefreshToken
		newRefreshToken.Token = accessData.RefreshToken
		if err := tx.Create(&newRefreshToken).Error; err != nil {
			tx.Rollback()
			return err
		}
		refreshTokenFK.Int64 = int64(newRefreshToken.ID)
		refreshTokenFK.Valid = true
	}

	// Put refresh token id, into OAuth2AccessData.refreshtoken fk
	oauthAccessData.RefreshTokenID = refreshTokenFK

	if err := tx.Create(oauthAccessData).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

// LoadAccess retrieves access data by token. osin.Client information MUST be loaded together.
// osin.AuthorizeData and osin.AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadAccess(token string) (*osin.AccessData, error) {
	oad := new(OAuth2AccessData)
	Db().Model(OAuth2AccessData{}).Where(&OAuth2AccessData{AccessToken: token}).Find(oad)
	if oad.AccessToken != token {
		return nil, errors.New("LoadAccess: AccessToken not found")
	}

	var ret osin.AccessData

	ret.CreatedAt = oad.CreatedAt
	ret.ExpiresIn = int32(oad.ExpiresIn)
	if ret.IsExpired() {
		return nil, errors.New("Access token expired")
	}

	if client, err := s.GetClient(strconv.FormatUint(oad.ClientID, 10)); err == nil {
		ret.Client = client
	} else {
		return nil, err
	}

	ret.AccessToken = token
	ret.Scope = oad.Scope
	ret.RedirectUri = oad.RedirectURI
	ret.UserData = oad.UserID

	if oad.RefreshTokenID.Valid {
		var refreshToken OAuth2RefreshToken
		if err := Db().First(&refreshToken, oad.RefreshTokenID.Int64).Error; err != nil {
			return nil, err
		}
		ret.RefreshToken = refreshToken.Token
	}

	return &ret, nil
}

// RemoveAccess revokes or deletes an osin.AccessData.
func (s *OAuth2Storage) RemoveAccess(token string) error {
	return Db().Where(&OAuth2AccessData{AccessToken: token}).Delete(OAuth2AccessData{}).Error
}

// LoadRefresh retrieves refresh osin.AccessData. osin.Client information MUST be loaded together.
// osin.AuthorizeData and osin.AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadRefresh(token string) (*osin.AccessData, error) {
	var pointedAccessData OAuth2AccessData
	var refreshToken OAuth2RefreshToken
	Db().Model(OAuth2RefreshToken{}).Where(&OAuth2RefreshToken{Token: token}).Find(&refreshToken)
	if refreshToken.Token != token {
		return nil, errors.New("Refresh token not found")
	}

	var refreshTokenNullInt64 sql.NullInt64
	refreshTokenNullInt64.Int64, refreshTokenNullInt64.Valid = int64(refreshToken.ID), true

	Db().Model(OAuth2AccessData{}).Where(&OAuth2AccessData{RefreshTokenID: refreshTokenNullInt64}).Find(&pointedAccessData)
	return s.LoadAccess(pointedAccessData.AccessToken)
}

// RemoveRefresh revokes or deletes refresh osin.AccessData.
func (s *OAuth2Storage) RemoveRefresh(token string) error {
	return Db().Where(&OAuth2RefreshToken{Token: token}).Delete(OAuth2RefreshToken{}).Error
}

// Implementing the osin.Client interface

// GetId returns the client ID
func (d *OAuth2Client) GetId() string {
	return strconv.FormatUint(d.ID, 10)
}

// GetSecret returns the client secret
func (d *OAuth2Client) GetSecret() string {
	return d.Secret
}

// GetRedirectUri returns the client redirect URI
func (d *OAuth2Client) GetRedirectUri() string {
	return d.RedirectURI
}

// GetUserData returns client UserData
func (d *OAuth2Client) GetUserData() interface{} {
	return d.UserID
}

// Additional methods

// RemoveClient removes the client by id (primary key)
func (s *OAuth2Storage) RemoveClient(id uint64) error {
	if id <= 0 {
		return errors.New("Invalid client id")
	}

	return Db().Where(&OAuth2Client{ID: id}).Delete(OAuth2Client{}).Error
}

// CreateClient creates a new OAuth2 Client
func (s *OAuth2Storage) CreateClient(c osin.Client) (*OAuth2Client, error) {
	client := OAuth2Client{
		RedirectURI: c.GetRedirectUri(),
		Secret:      c.GetSecret(),
		UserID:      c.GetUserData().(uint64),
	}

	if err := Db().Create(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

// UpdateClient update client with id c.GetId()
func (s *OAuth2Storage) UpdateClient(c osin.Client) (*OAuth2Client, error) {
	var numericID uint64
	var err error
	if numericID, err = strconv.ParseUint(c.GetId(), 10, 64); err != nil {
		return nil, fmt.Errorf("invalid client_id: %s", c.GetId())
	}

	client := OAuth2Client{
		ID:          numericID,
		RedirectURI: c.GetRedirectUri(),
		Secret:      c.GetSecret(),
		UserID:      c.GetUserData().(uint64),
	}

	fmt.Println(client)

	if err := Db().Save(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

// HandleLoginPage is an helper used by the OAuth2 authentication process to login the user (if it's not logged)
// and to show a basic login form that redirect to the authorization endpoint
func HandleLoginPage(ar *osin.AuthorizeRequest, c *echo.Context) (*User, error) {
	r := c.Request()
	r.ParseForm()
	if r.Method == "POST" {
		if user, err := Login(r.Form.Get("login"), r.Form.Get("password")); err == nil { // succcessful logged in
			return user, nil
		} else {
			return nil, err
		}
	}

	var buffer bytes.Buffer
	buffer.WriteString("<html><body>")
	buffer.WriteString(fmt.Sprintf("LOGIN %s<br />", ar.Client.GetId()))
	buffer.WriteString(fmt.Sprintf("<form action=\"/authorize?response_type=%s&client_id=%s&state=%s&redirect_uri=%s\" method=\"POST\">",
		ar.Type, ar.Client.GetId(), ar.State, url.QueryEscape(ar.RedirectUri)))

	buffer.WriteString("Login: <input type=\"text\" name=\"login\" /><br/>")
	buffer.WriteString("Password: <input type=\"password\" name=\"password\" /><br/>")
	buffer.WriteString("<input type=\"submit\"/>")
	buffer.WriteString("</form></body></html>")
	c.HTML(http.StatusOK, buffer.String())
	return nil, errors.New("Please login")
}

// DownloadAccessToken is an helper used by the OAuth2 basic authentication process. It downloads the access token
func DownloadAccessToken(url string, auth *osin.BasicAuth, output map[string]interface{}) error {
	// download access token
	preq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if auth != nil {
		preq.SetBasicAuth(auth.Username, auth.Password)
	}

	pclient := &http.Client{}
	presp, err := pclient.Do(preq)
	if err != nil {
		return err
	}

	if presp.StatusCode != 200 {
		return errors.New("Invalid status code")
	}

	jdec := json.NewDecoder(presp.Body)
	err = jdec.Decode(&output)
	return err
}
