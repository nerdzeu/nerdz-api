package orm

import (
    "net/url"
    "crypto/md5"
    "time"
    "strings"
    "io"
    "fmt"
    "errors"
)

// JSON return value

type UserInfo struct {
    Id       int64
    IsOnline bool
    Nation   string
    Timezone string
    Username string
    Name     string
    Surname  string
    Gender   bool
    Birthday time.Time
    Gravatar url.URL
}

func getGravatar(email string) url.URL {

    m := md5.New()
    io.WriteString(m, strings.ToLower(email))

    return url.URL{
        Scheme: "https",
        Host: "www.gravatar.com",
        Path: "/avatar/" + fmt.Sprintf("%x", m.Sum(nil)) }

}

func (*User) GetInfo(id int64) (*UserInfo, error) {
    var user User
    var profile Profile

    db.First(&user, id)
    fmt.Printf("%+v",user)
    db.Find(&profile, id)

    if user.Counter != id || profile.Id != id {
        return nil, errors.New("Invalid id")
    }

    return &UserInfo {
        Id: id,
        IsOnline: user.Viewonline && user.Last.Add(time.Duration(5)*time.Minute).After(time.Now()),
        Nation: user.Lang,
        Timezone: user.Timezone,
        Name: user.Name,
        Surname: user.Surname,
        Gender: user.Gender,
        Birthday: user.BirthDate,
        Gravatar: getGravatar(user.Email) }, nil
}
