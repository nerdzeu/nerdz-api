package nerdz

import (
	"flag"
	"fmt"
	"log"
	"os"

	// Blank import required to get Postgresql working
	_ "github.com/lib/pq"
)

var db *Database

// This is the first method called. Parse the configuration file, populate the environment values and create the connection to the db
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

	var err error

	if err = initConfiguration(file); err != nil {
		panic(fmt.Sprintf("[!] %v\n", err))
	}

	var connectionString string
	if connectionString, err = Configuration.ConnectionString(); err != nil {
		panic(err.Error())
	}

	if db, err = Connect(connectionString); err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}

	// Ping database to effectively check database connection
	if err := db.DB().Ping(); err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}

	logger := log.New(os.Stdout, "query-logger: ", log.LUTC)
	db.Log(logger)
}

// Db returns the *database
func Db() *Database {
	return db
}
