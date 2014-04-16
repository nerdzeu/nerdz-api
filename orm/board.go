package orm

import "net/url"

type Info struct {
    Id  int64
    Owner *User
    Followers []*User
    Name string
    Website *url.URL
    Image *url.URL
}

type Board interface {
    GetInfo() *Info
}
