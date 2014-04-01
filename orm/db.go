package orm

import (
        "github.com/jinzhu/gorm"
        _ "github.com/lib/pq"
        "github.com/nerdzeu/nerdz-api/utils"
        "fmt"
        "flag"
)

var DB gorm.DB

func init() {
    flag.Parse()

    args := flag.Args()

    if len(args) != 1 {
        panic(fmt.Sprintln("Configuration file is required"))
    }

    connStr, err := utils.Parse(args[0])

    if err != nil {
        panic(fmt.Sprintf("[!] %v\n", err))
    }

    DB, err = gorm.Open("postgres", connStr)

    if err != nil {
        panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
    }
}
