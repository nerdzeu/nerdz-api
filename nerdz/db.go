package nerdz

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db gorm.DB

// Callback function: invoked after the creation/update of an object. To populate its default fields
func updateFields(scope *gorm.Scope) {
	if !scope.HasError() {
		newScope := scope.New(scope.Value)
		newScope.Search.TableName = scope.TableName()
		gorm.Query(newScope)
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
func InitDB() {
	connectionString, errConn := Configuration.ConnectionString()
	if errConn != nil {
		panic(errConn.Error())
	}

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}

	if Configuration.EnableLog == 1 {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

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
