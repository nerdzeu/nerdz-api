package nerdz

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"reflect"
)

var db gorm.DB

// Callback function: invoked after the creation/update of an object. To populate its default fields
func updateFields(scope *gorm.Scope) {
	if !scope.HasError() {
		newScope := scope.New(scope.Value)
		newScope.Search.TableName = scope.TableName()
		if scope.PrimaryKey() != "" {
			gorm.Query(newScope)
		} else {
			// TODO: find a way to populate fields of scope.Value selecting * matching on every fields
		}
		scope = newScope
	}
}

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

	if err = InitConfiguration(file); err != nil {
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

	db.LogMode(Configuration.EnableLog)

	// Remove default useless gorm callbacks for the nerdz-db architecture
	// update
	db.Callback().Update().Remove("gorm:save_before_associations")
	db.Callback().Update().Remove("gorm:update_time_stamp_when_update")
	db.Callback().Update().Remove("gorm:save_after_associations")
	// create
	db.Callback().Create().Remove("gorm:save_before_associations")
	db.Callback().Create().Remove("gorm:update_time_stamp_when_update")
	db.Callback().Create().Remove("gorm:save_after_associations")

	// Add after update/create callback to populate the struct after and update/create query
	db.Callback().Update().After("gorm:create").Register("nerdz-api:update_fields", updateFields)
	db.Callback().Create().After("gorm:create").Register("nerdz-api:update_fields", updateFields)

	// Clear field values after delete
	db.Callback().Delete().Register("nerdz-api:clear_fields", clearFields)
}
