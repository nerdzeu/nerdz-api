package orm

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nerdzeu/gorm"
	"github.com/nerdzeu/nerdz-api/utils"
)

var db gorm.DB

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

	db, err = gorm.Open("postgres", connStr)

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}
}
