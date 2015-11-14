package nerdz_test

import (
	"github.com/RangelReale/osin"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"reflect"
	"testing"
	/*
		"database/sql"


		"log"
		"os"
		"testing"
		"time" */)

var store *nerdz.OAuth2Storage
var client1, client2 *nerdz.OAuth2Client

func TestCreateApplication(t *testing.T) {
	create := &osin.DefaultClient{
		Secret:      "secret 1",
		RedirectUri: "http://localhost/",
		UserData:    me.Counter,
	}

	if client1, err = store.CreateClient(create); err != nil {
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

	if client2, err = store.CreateClient(create2); err != nil {
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
	authorize.Scope = "update_profile notifications public_messages"

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
		Scope:       "public_messages",
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
		Scope:         "public_messages",
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
		Scope:         "notifications",
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

	if !reflect.DeepEqual(access, result) {
		t.Errorf("access and result shoud be equal, but are different:\n%v\n%v\n", access, result)
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

/*
func TestRefreshOperations(t *testing.T) {
	client := &osin.DefaultClient{"4", "secret", "http://localhost/", ""}
	type test struct {
		access *osin.AccessData
	}

	for k, c := range []*test{
		&test{
			access: &osin.AccessData{
				Client: client,
				AuthorizeData: &osin.AuthorizeData{
					Client:      client,
					Code:        uuid.New(),
					ExpiresIn:   int32(60),
					Scope:       "scope",
					RedirectUri: "http://localhost/",
					State:       "state",
					CreatedAt:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
					UserData:    userDataMock,
				},
				AccessData:   nil,
				AccessToken:  uuid.New(),
				RefreshToken: uuid.New(),
				ExpiresIn:    int32(60),
				Scope:        "scope",
				RedirectUri:  "https://localhost/",
				CreatedAt:    time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				UserData:     userDataMock,
			},
		},
	} {
		createClient(t, store, client)
		require.Nil(t, store.SaveAuthorize(c.access.AuthorizeData), "Case %d", k)
		require.Nil(t, store.SaveAccess(c.access), "Case %d", k)

		result, err := store.LoadRefresh(c.access.RefreshToken)
		require.Nil(t, err)
		require.True(t, reflect.DeepEqual(c.access, result), "Case %d", k)

		require.Nil(t, store.RemoveRefresh(c.access.RefreshToken))
		_, err = store.LoadRefresh(c.access.RefreshToken)

		require.NotNil(t, err, "Case %d", k)
		require.Nil(t, store.RemoveAccess(c.access.AccessToken), "Case %d", k)
		require.Nil(t, store.SaveAccess(c.access), "Case %d", k)

		_, err = store.LoadRefresh(c.access.RefreshToken)
		require.Nil(t, err, "Case %d", k)

		require.Nil(t, store.RemoveAccess(c.access.AccessToken), "Case %d", k)
		_, err = store.LoadRefresh(c.access.RefreshToken)
		require.NotNil(t, err, "Case %d", k)

	}
	removeClient(t, store, client)
}

func getClient(t *testing.T, store storage.Storage, set osin.Client) {
	client, err := store.GetClient(set.GetId())
	require.Nil(t, err)
	require.EqualValues(t, set, client)
}

func createClient(t *testing.T, store storage.Storage, set osin.Client) {
	require.Nil(t, store.CreateClient(set))
}

func updateClient(t *testing.T, store storage.Storage, set osin.Client) {
	require.Nil(t, store.UpdateClient(set))
}

func removeClient(t *testing.T, store storage.Storage, set osin.Client) {
	require.Nil(t, store.RemoveClient(set.GetId()))
}
*/
