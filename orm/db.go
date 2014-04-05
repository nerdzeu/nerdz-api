package orm

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nerdzeu/gorm"
	"github.com/nerdzeu/nerdz-api/utils"
	"os"
)

var db gorm.DB

func init() {
	flag.Parse()

	args := flag.Args()
	envVar := os.Getenv("CONF_FILE")

	var file string

	if len(args) == 1 {
		file = args[0]
	} else if envVar != "" {
		file = envVar
	} else {
		panic(fmt.Sprintln("Configuration file is required.\nUse: CONF_FILE environment variable or cli args"))
	}

	connStr, err := utils.Parse(file)

	if err != nil {
		panic(fmt.Sprintf("[!] %v\n", err))
	}

	db, err = gorm.Open("postgres", connStr)

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}
}
