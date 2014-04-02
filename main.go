package main

import (
        "fmt"
        // Invoke init(): parse configuration file, connect to database and create DB var
        "github.com/nerdzeu/nerdz-api/orm"
)

func main() {
    var info *orm.Info
    var user orm.User

    info, err := user.GetInfo(1);

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("%+v\n",info.Gravatar);
}
