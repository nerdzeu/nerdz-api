package orm_test

import (
        "testing"
        "github.com/nerdzeu/nerdz-api/orm"
        "fmt"
       )

var user orm.User

func init() {
    err := user.New(1)
    if err != nil {
        panic(fmt.Sprintf("No error should happen when create existing user, but got: %+v", err))
    }
}

func TestGetContactInfo(t *testing.T) {
    info := user.GetContactInfo()
    if info == nil {
        t.Error("null info")
    }

    fmt.Printf("%v\n",info);
}

func TestGetPersonalInfo(t *testing.T) {
    info := user.GetContactInfo()
    if info == nil {
        t.Error("null info")
    }

    fmt.Printf("%v\n",info);
}
