package nerdz

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/jinzhu/gorm"
	// Blank import required to get Postgresql working
	_ "github.com/lib/pq"
)

var db gorm.DB

// Db export the private var Db outside the nerdz package
func Db() *gorm.DB {
	return &db
}

// Callback function: invoked after the creation/update of an object. To populate its default fields
/*
func updateFields(scope *gorm.Scope) {
	if !scope.HasError() {
		newScope := scope.New(scope.Value)
		newScope.Search = newScope.Search.Table(scope.TableName())

		if scope.PrimaryKey() != "" {
			gorm.Query(newScope)
		} else {
			// TODO: find a way to populate fields of scope.Value selecting * matching on every fields

		}
		scope = newScope
	}
}
*/
// Now gorm hash "gorm:force_reload_after_create" when a field has a default value

// Callback function: invoked after the deletion of an object. To clear its fields
func clearFields(scope *gorm.Scope) {
	if !scope.HasError() {
		scope.Value = reflect.New(reflect.TypeOf(scope.Value))
	}
}

// This is the first methdo to be called. Parse the configuration file, populate the environment values and create the connection to the db
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

	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}

	// Ping database to effectively check database connection
	if err := db.DB().Ping(); err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}

	Db().LogMode(Configuration.EnableLog)

	// Remove default useless gorm callbacks for the nerdz-db architecture
	// update
	Db().Callback().Update().Remove("gorm:save_before_associations")
	Db().Callback().Update().Remove("gorm:update_time_stamp_when_update")
	Db().Callback().Update().Remove("gorm:save_after_associations")
	// create
	Db().Callback().Create().Remove("gorm:save_before_associations")
	Db().Callback().Create().Remove("gorm:update_time_stamp_when_create")
	Db().Callback().Create().Remove("gorm:save_after_associations")

	// Add after update/create callback to populate the struct after and update/create query
	/*
		Db().Callback().Update().After("gorm:create").Register("nerdz-api:update_fields", updateFields)
		Db().Callback().Create().After("gorm:create").Register("nerdz-api:update_fields", updateFields)
	*/

	// Clear field values after delete
	Db().Callback().Delete().Register("nerdz-api:clear_fields", clearFields)
}
