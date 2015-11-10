package nerdz_test

import (
	"github.com/RangelReale/osin"
	"github.com/nerdzeu/nerdz-api/nerdz"
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
}

/*
func TestAuthorizeOperations(t *testing.T) {
	client := &osin.DefaultClient{"2", "secret 1", "http://localhost/", ""}
	createClient(t, store, client)

	for k, authorize := range []*osin.AuthorizeData{
		&osin.AuthorizeData{
			Client:      client,
			Code:        uuid.New(),
			ExpiresIn:   int32(60),
			Scope:       "scope",
			RedirectUri: "http://localhost/",
			State:       "state",
			// FIXME this should be time.Now(), but an upstream ( https://github.com/lib/pq/issues/329 ) issue prevents this.
			CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			UserData:  userDataMock,
		},
	} {
		// Test save
		require.Nil(t, store.SaveAuthorize(authorize))

		// Test fetch
		result, err := store.LoadAuthorize(authorize.Code)
		require.Nil(t, err)
		require.True(t, reflect.DeepEqual(authorize, result), "Case: %d\n%v\n\n%v", k, authorize, result)

		// Test remove
		require.Nil(t, store.RemoveAuthorize(authorize.Code))
		_, err = store.LoadAuthorize(authorize.Code)
		require.NotNil(t, err)
	}

	removeClient(t, store, client)
}
*/
/*
func TestStoreFailsOnInvalidUserData(t *testing.T) {
	client := &osin.DefaultClient{"3", "secret", "http://localhost/", ""}
	authorize := &osin.AuthorizeData{
		Client:      client,
		Code:        uuid.New(),
		ExpiresIn:   int32(60),
		Scope:       "scope",
		RedirectUri: "http://localhost/",
		State:       "state",
		CreatedAt:   time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UserData:    struct{ foo string }{"bar"},
	}
	access := &osin.AccessData{
		Client:        client,
		AuthorizeData: authorize,
		AccessData:    nil,
		AccessToken:   uuid.New(),
		RefreshToken:  uuid.New(),
		ExpiresIn:     int32(60),
		Scope:         "scope",
		RedirectUri:   "https://localhost/",
		CreatedAt:     time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UserData:      struct{ foo string }{"bar"},
	}
	assert.NotNil(t, store.SaveAuthorize(authorize))
	assert.NotNil(t, store.SaveAccess(access))
}

func TestAccessOperations(t *testing.T) {
	client := &osin.DefaultClient{"3", "secret", "http://localhost/", ""}
	authorize := &osin.AuthorizeData{
		Client:      client,
		Code:        uuid.New(),
		ExpiresIn:   int32(60),
		Scope:       "scope",
		RedirectUri: "http://localhost/",
		State:       "state",
		// FIXME this should be time.Now(), but an upstream ( https://github.com/lib/pq/issues/329 ) issue prevents this.
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UserData:  userDataMock,
	}
	nestedAccess := &osin.AccessData{
		Client:        client,
		AuthorizeData: authorize,
		AccessData:    nil,
		AccessToken:   uuid.New(),
		RefreshToken:  uuid.New(),
		ExpiresIn:     int32(60),
		Scope:         "scope",
		RedirectUri:   "https://localhost/",
		// FIXME this should be time.Now(), but an upstream ( https://github.com/lib/pq/issues/329 ) issue prevents this.
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UserData:  userDataMock,
	}
	access := &osin.AccessData{
		Client:        client,
		AuthorizeData: authorize,
		AccessData:    nestedAccess,
		AccessToken:   uuid.New(),
		RefreshToken:  uuid.New(),
		ExpiresIn:     int32(60),
		Scope:         "scope",
		RedirectUri:   "https://localhost/",
		// FIXME this should be time.Now(), but an upstream ( https://github.com/lib/pq/issues/329 ) issue prevents this.
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UserData:  userDataMock,
	}

	createClient(t, store, client)
	require.Nil(t, store.SaveAuthorize(authorize))
	require.Nil(t, store.SaveAccess(nestedAccess))
	require.Nil(t, store.SaveAccess(access))

	result, err := store.LoadAccess(access.AccessToken)
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(access, result))

	require.Nil(t, store.RemoveAccess(access.AccessToken))
	_, err = store.LoadAccess(access.AccessToken)
	require.NotNil(t, err)
	require.Nil(t, store.RemoveAuthorize(authorize.Code))

	removeClient(t, store, client)
}

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
