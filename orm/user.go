package orm

import (
    "net/url"
    "time"
    "errors"
)

// JSON return value

type Info struct {
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

func (*User) GetInfo(id int64) (*Info, error) {
    var user User
    var profile Profile

    db.First(&user, id)
    db.Find(&profile, id)

    if user.Counter != id || profile.Counter != id {
        return nil, errors.New("Invalid id")
    }

    return &Info {
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
