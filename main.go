package main

import (
//        "github.com/jinzhu/gorm"
//        _ "github.com/lib/pq"
        "github.com/nerdzeu/nerdz-api/utils"
        "fmt"
        "flag"
)


func main() {
    flag.Parse()

    args := flag.Args()

    if len(args) != 1 {
        fmt.Println("Configuration file is required")
        return
    }

    connStr, err := utils.Parse(args[0]);

    if err != nil {
        fmt.Printf("[!] %s\n", err.Error())
        return
    }

    fmt.Printf("Connection string:\n%s\n",connStr);

//    db, err := gorm.Open("postgres", connStr);

    return
}
