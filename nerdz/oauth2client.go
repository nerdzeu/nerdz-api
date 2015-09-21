package nerdz

// Implements osin.Client interface

// GetId returns the client ID
func (d *OAuth2Client) GetId() string {
	return d.ID
}

// GetSecret returns the client secret
func (d *OAuth2Client) GetSecret() string {
	return d.Secret
}

// GetRedirectUri returns the client redirect URI
func (d *OAuth2Client) GetRedirectUri() string {
	return d.RedirectUri
}

// GetUserData returns client UserData
func (d *OAuth2Client) GetUserData() interface{} {
	return d.UserData
}
