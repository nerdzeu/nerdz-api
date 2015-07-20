package nerdz

import (
	"errors"
	"github.com/RangelReale/osin"
)

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
	client := new(OAuth2Client)
	Db().First(client, id)
	if client.GetId() != id {
		return nil, errors.New("Client not found")
	}
	return client, nil
}

// SaveAuthorize saves authorize data.
func (s *OAuth2Storage) SaveAuthorize(data *osin.AuthorizeData) error {
	d := &OAuth2AuthorizeData{
		ClientID:    data.Client.GetId(),
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		RedirectUri: data.RedirectUri,
		Scope:       data.Scope,
		State:       data.State,
		UserData:    data.UserData.([]byte)}

	return Db().Create(&d).Error
}

// LoadAuthorize looks up osin.AuthorizeData by a code.
// osin.Client information MUST be loaded together.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	authorize := new(OAuth2AuthorizeData)
	Db().First(authorize, code)
	if code != authorize.Code {
		return nil, errors.New("Authorization data not found")
	}

	authData := &osin.AuthorizeData{
		Code:        code,
		ExpiresIn:   authorize.ExpiresIn,
		RedirectUri: authorize.RedirectUri,
		Scope:       authorize.Scope,
		State:       authorize.State,
		UserData:    authorize.UserData}

	if client, err := s.GetClient(authorize.ClientID); err != nil {
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
	oauthAccessData := OAuth2AccessData{
		AccessToken:     accessData.AccessToken,
		AuthorizeDataID: accessData.AuthorizeData.Code,
		ClientID:        accessData.Client.GetId(),
		CreatedAt:       accessData.CreatedAt,
		ExpiresIn:       accessData.ExpiresIn,
		AccessDataID:    accessData.AccessData.AccessToken,
		RedirectUri:     accessData.RedirectUri,
		RefreshToken:    accessData.RefreshToken,
		Scope:           accessData.Scope,
		UserData:        accessData.UserData.([]byte)}
	return Db().Create(oauthAccessData).Error
}

// LoadAccess retrieves access data by token. osin.Client information MUST be loaded together.
// osin.AuthorizeData and osin.AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadAccess(token string) (*osin.AccessData, error) {
	oad := new(OAuth2AccessData)
	Db().First(&oad, token)
	if oad.AccessToken != token {
		return nil, errors.New("LoadAccess: AccessToken not found")
	}

	var ret osin.AccessData
	ret.ExpiresIn = oad.ExpiresIn
	if ret.IsExpired() {
		return nil, errors.New("Access token expired")
	}
	ret.Scope = oad.Scope
	ret.RedirectUri = oad.RedirectUri
	ret.CreatedAt = oad.CreatedAt
	ret.UserData = oad.UserData

	if client, err := s.GetClient(oad.ClientID); err != nil {
		ret.Client = client
	} else {
		return nil, errors.New("LoadAccess: Client not found")
	}

	var authData OAuth2AuthorizeData
	if err := Db().First(&authData, oad.AuthorizeDataID).Error; err != nil {
		return nil, err
	}
	ret.AuthorizeData.Code = authData.Code
	ret.AuthorizeData.CreatedAt = authData.CreatedAt
	ret.AuthorizeData.ExpiresIn = authData.ExpiresIn
	ret.AuthorizeData.RedirectUri = authData.RedirectUri
	ret.AuthorizeData.Scope = authData.Scope
	ret.AuthorizeData.State = authData.State
	ret.AuthorizeData.UserData = authData.UserData
	ret.AuthorizeData.Client = ret.Client

	if oad.RefreshToken != "" {
		var e error
		ret.AccessData, e = s.LoadAccess(oad.RefreshToken)
		if e != nil {
			return nil, e
		}
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
	Db().First(&pointedAccessData, &OAuth2AccessData{RefreshToken: token})
	if pointedAccessData.RefreshToken != token {
		return nil, errors.New("AccessData not found")
	}
	return s.LoadAccess(pointedAccessData.AccessToken)
}

// RemoveRefresh revokes or deletes refresh osin.AccessData.
func (s *OAuth2Storage) RemoveRefresh(token string) error {
	return Db().Where(&OAuth2AccessData{RefreshToken: token}).Delete(OAuth2AccessData{}).Error
}
