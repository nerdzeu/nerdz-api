package main

import (
        "fmt"
        // Invoke init(): parse configuration file, connect to database and create DB var
        _ "github.com/nerdzeu/nerdz-api/orm"
)

func main() {
    fmt.Println("It works");
}
