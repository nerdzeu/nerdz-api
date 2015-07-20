package nerdz

/*

//TODO: implement the osin.Storage interface, create tables on db and create gorm models

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
    if(client.ID() != id) {
        return nil, errors.New("Client not found")
    }
    return client, nil
}

// SaveAuthorize saves authorize data.
func (s *OAuth2Storage) SaveAuthorize(*osin.AuthorizeData) error {
//    Db().Create(
}

// LoadAuthorize looks up osin.AuthorizeData by a code.
// osin.Client information MUST be loaded together.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
}

// RemoveAuthorize revokes or deletes the authorization code.
func (s *OAuth2Storage) RemoveAuthorize(code string) error {
}

// SaveAccess writes osin.AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (s *OAuth2Storage) SaveAccess(*osin.AccessData) error {
}

// LoadAccess retrieves access data by token. osin.Client information MUST be loaded together.
// osin.AuthorizeData and osin.AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadAccess(token string) (*osin.AccessData, error) {
}

// RemoveAccess revokes or deletes an osin.AccessData.
func (s *OAuth2Storage) RemoveAccess(token string) error {
}

// LoadRefresh retrieves refresh osin.AccessData. osin.Client information MUST be loaded together.
// osin.AuthorizeData and osin.AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *OAuth2Storage) LoadRefresh(token string) (*osin.AccessData, error) {
}

// RemoveRefresh revokes or deletes refresh osin.AccessData.
func (s *OAuth2Storage) RemoveRefresh(token string) error {
} */
