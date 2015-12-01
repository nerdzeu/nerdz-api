package nerdz_test

import (
	"github.com/RangelReale/osin"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"reflect"
	"testing"
)

var store *nerdz.OAuth2Storage
var client1, client2 *nerdz.OAuth2Client

func TestCreateApplication(t *testing.T) {
	create := &osin.DefaultClient{
		Secret:      "secret 1",
		RedirectUri: "http://localhost/",
		UserData:    me.Counter,
	}

	if client1, err = store.CreateClient(create, "Application 1"); err != nil {
		t.Errorf("Unable to create application client1: %s\n", err.Error())
	}

	update := &osin.DefaultClient{
		Id:          client1.GetId(),
		Secret:      client1.GetSecret(),
		UserData:    client1.GetUserData(),
		RedirectUri: "http://www.nerdz.eu",
	}

	if client1, err = store.UpdateClient(update); err != nil {
		t.Errorf("Unable to update application client1, redirectURI: %s\n", err.Error())
	}

	create2 := &osin.DefaultClient{
		Secret:      "secret 2",
		RedirectUri: "http://localhost/",
		UserData:    me.Counter,
	}

	if client2, err = store.CreateClient(create2, "Application 2"); err != nil {
		t.Errorf("Unable to create application client2: %s\n", err.Error())
	}
}

func TestAuthorizeOperationsAndGetCient(t *testing.T) {
	var client osin.Client
	if client, err = store.GetClient("1"); err != nil {
		t.Error(err.Error())
	}

	authorizeInvlid := &osin.AuthorizeData{
		Client:      client,
		Code:        "auth code 1",
		ExpiresIn:   int32(60),
		Scope:       "invalid",
		RedirectUri: "http://localhost/",
		State:       "state",
		// CreatedAt field is automatically filled by the db
		UserData: me.Counter,
	}

	if err = store.SaveAuthorize(authorizeInvlid); err == nil {
		t.Errorf("This authorization should fail")
	}

	authorize := authorizeInvlid
	authorize.Scope = "profile:read notifications:read,write profile_messages:write"

	if err = store.SaveAuthorize(authorize); err != nil {
		t.Errorf("Not should work, but got: %s\n", err.Error())
	}

	// Test fetch
	var result *osin.AuthorizeData
	if result, err = store.LoadAuthorize(authorize.Code); err != nil {
		t.Errorf("Unable to load AuthorizeData with code %s, got error: %s", authorize.Code, err.Error())
	}

	// Since createdAt is created by the dbms
	authorize.CreatedAt = result.CreatedAt
	if !reflect.DeepEqual(*authorize, *result) {
		t.Errorf("authorize and result are different, %v\n, \n%v", *authorize, *result)
	}

	// Test remove
	if err = store.RemoveAuthorize(authorize.Code); err != nil {
		t.Errorf("RemoveAuthozire should work, but got: %s\n", err.Error())
	}

	// check if it was really removed
	if _, err = store.LoadAuthorize(authorize.Code); err == nil {
		t.Errorf("Authorization not removed")
	}
}

// there's no need to check this in our implementations, since the dbms have
// foreign key constrint un user data (user data = User id = foreign key to users)
func TestStoreFailsOnInvalidUserData(t *testing.T) {
}

func TestAccessOperations(t *testing.T) {
	var err error
	authorize := &osin.AuthorizeData{
		Client:      client2,
		Code:        "code lel",
		ExpiresIn:   int32(60),
		Scope:       "project_messages:read",
		RedirectUri: "http://localhost/",
		State:       "state",
		UserData:    me.Counter,
	}
	nestedAccess := &osin.AccessData{
		Client:        client2,
		AuthorizeData: authorize,
		AccessData:    nil,
		AccessToken:   "new random access token",
		RefreshToken:  "new random refresh token",
		ExpiresIn:     int32(60),
		Scope:         "project_messages:write",
		RedirectUri:   "https://localhost/",
		UserData:      me.Counter,
	}
	access := &osin.AccessData{
		Client:        client2,
		AuthorizeData: authorize,
		AccessData:    nestedAccess,
		AccessToken:   "other new random access token",
		RefreshToken:  "other new random refresh token",
		ExpiresIn:     int32(60),
		Scope:         "notifications:read",
		RedirectUri:   "https://localhost/",
		UserData:      me.Counter,
	}

	if err = store.SaveAuthorize(authorize); err != nil {
		t.Errorf("SaveAuthorize should work but got: %s\n", err.Error())
	}

	if err = store.SaveAccess(nestedAccess); err != nil {
		t.Errorf("SaveAccess should work but got: %s\n", err.Error())
	}

	if err = store.SaveAccess(access); err != nil {
		t.Errorf("SaveAccess should work but got: %s\n", err.Error())
	}

	var result *osin.AccessData
	if result, err = store.LoadAccess(access.AccessToken); err != nil {
		t.Errorf("LoadAccess should work but got: %s\n", err.Error())
	}

	// Since createdAt is created by the dbms
	access.CreatedAt = result.CreatedAt
	// AccessData and Authorize data are optional, and thus not filled by LoadAccess
	access.AccessData = nil
	access.AuthorizeData = nil
	if !reflect.DeepEqual(*access, *result) {
		t.Errorf("access and result shoud be equal, but are different:\n%v\n%v\n", *access, *result)
	}

	if err = store.RemoveAccess(access.AccessToken); err != nil {
		t.Errorf("RemoveAccess should work but got: %s\n", err.Error())
	}

	if _, err = store.LoadAccess(access.AccessToken); err == nil {
		t.Errorf("LoadAccess should fail, but it worked")
	}

	if err = store.RemoveAuthorize(authorize.Code); err != nil {
		t.Errorf("RemoveAuthozire should work but got: %s\n", err.Error())
	}

}

func TestRefreshOperations(t *testing.T) {
	var err error

	access := &osin.AccessData{
		Client: client2,
		AuthorizeData: &osin.AuthorizeData{
			Client:      client2,
			Code:        "nice code",
			ExpiresIn:   int32(60),
			Scope:       "profile_messages:write",
			RedirectUri: "http://localhost/",
			State:       "state",
			UserData:    me.Counter,
		},
		AccessData:   nil,
		AccessToken:  "nice access token",
		RefreshToken: "nice refresh token",
		ExpiresIn:    int32(60),
		Scope:        "notifications:read",
		RedirectUri:  "https://localhost/",
		UserData:     me.Counter,
	}

	if err = store.SaveAuthorize(access.AuthorizeData); err != nil {
		t.Errorf("%s", err.Error())
	}

	if err = store.SaveAccess(access); err != nil {
		t.Errorf("%s", err.Error())
	}

	var result *osin.AccessData
	if result, err = store.LoadRefresh(access.RefreshToken); err != nil {
		t.Errorf("%s", err.Error())
	}

	access.CreatedAt = result.CreatedAt
	backAuthorize := access.AuthorizeData
	backAccesData := access.AccessData
	access.AuthorizeData = nil
	access.AccessData = nil
	if !reflect.DeepEqual(*access, *result) {
		t.Errorf("access and result are different, %v\n, \n%v", *access, *result)
	}

	access.AuthorizeData = backAuthorize
	access.AccessData = backAccesData

	if err = store.RemoveRefresh(access.RefreshToken); err != nil {
		t.Errorf("%s", err.Error())
	}

	if _, err = store.LoadRefresh(access.RefreshToken); err == nil {
		t.Errorf("refresh token not removed :(")
	}

	if err = store.RemoveAccess(access.AccessToken); err != nil {
		t.Errorf("%s", err.Error())
	}

	if err = store.SaveAccess(access); err != nil {
		t.Errorf("%s", err.Error())
	}

	if _, err = store.LoadRefresh(access.RefreshToken); err != nil {
		t.Errorf("%s", err.Error())
	}

	if store.RemoveAccess(access.AccessToken); err != nil {
		t.Errorf("%s", err.Error())
	}

	if _, err = store.LoadRefresh(access.RefreshToken); err == nil {
		t.Errorf("Previous RemoveAccess do not deleted related RefreshToken")
	}
}
