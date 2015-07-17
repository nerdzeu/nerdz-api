package oauth2

//TODO: implement the osin.Storage interface, create tables on db and create gorm models

import "github.com/RangelReale/osin"

// NerdzStorage implements osin.Storage interface
type NerdzStorage struct {
}

// Clone the storage if needed. For example, using mgo, you can clone the session with session.Clone
// to avoid concurrent access problems.
// This is to avoid cloning the connection at each method access.
// Can return itself if not a problem.
func (s *NerdzStorage) Clone() osin.Storage {
	return s
}

// Close the resources the Storage potentially holds (using Clone for example)
func (s *NerdzStorage) Close() {
}

// GetClient loads the client by id (client_id)
func (s *NerdzStorage) GetClient(id string) (osin.Client, error) {
}

// SaveAuthorize saves authorize data.
func (s *NerdzStorage) SaveAuthorize(*osin.AuthorizeData) error {
}

// LoadAuthorize looks up osin.AuthorizeData by a code.
// osin.Client information MUST be loaded together.
// Optionally can return error if expired.
func (s *NerdzStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
}

// RemoveAuthorize revokes or deletes the authorization code.
func (s *NerdzStorage) RemoveAuthorize(code string) error {
}

// SaveAccess writes osin.AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (s *NerdzStorage) SaveAccess(*osin.AccessData) error {
}

// LoadAccess retrieves access data by token. osin.Client information MUST be loaded together.
// osin.AuthorizeData and osin.AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *NerdzStorage) LoadAccess(token string) (*osin.AccessData, error) {
}

// RemoveAccess revokes or deletes an osin.AccessData.
func (s *NerdzStorage) RemoveAccess(token string) error {
}

// LoadRefresh retrieves refresh osin.AccessData. osin.Client information MUST be loaded together.
// osin.AuthorizeData and osin.AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (s *NerdzStorage) LoadRefresh(token string) (*osin.AccessData, error) {
}

// RemoveRefresh revokes or deletes refresh osin.AccessData.
func (s *NerdzStorage) RemoveRefresh(token string) error {
}
