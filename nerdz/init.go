/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package nerdz

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/galeone/igor"
)

var db *igor.Database

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

	if db, err = igor.Connect(connectionString); err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}

	logger := log.New(os.Stdout, "query-logger: ", log.LUTC)
	Db().Log(logger)
}

// Db returns the *igor.Database
func Db() *igor.Database {
	return db
}
