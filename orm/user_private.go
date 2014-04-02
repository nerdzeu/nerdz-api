package orm

import (
    "net/url"
    "crypto/md5"
    "strings"
    "io"
    "fmt"
)


func getGravatar(email string) url.URL {

    m := md5.New()
    io.WriteString(m, strings.ToLower(email))

    return url.URL{
        Scheme: "https",
        Host: "www.gravatar.com",
        Path: "/avatar/" + fmt.Sprintf("%x", m.Sum(nil)) }

}
