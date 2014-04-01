package orm

import (
    //"github.com/jinzhu/gorm"
    //"encoding/json"
    "net/url"
    "time"
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
/*
func GetUserInfo(id int64) UserInfo {
    info UserInfo
    return infoa
}*/
