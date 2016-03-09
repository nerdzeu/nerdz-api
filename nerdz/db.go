package nerdz

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
	// Blank import required to get PostgreSQL working
	_ "github.com/lib/pq"
	"github.com/nerdzeu/nerdz-api/utils"
)

// Reserved keyworkds that shouldn't be used as column name or other identifiers
// result of: select '"' || word::text || '",' from pg_get_keywords() where catcode = 'R' OR catdesc like '%cannot%';
// We will check if the current identifier is in that list. If it is, it will be placed between quotes
var reservedKeywords = []string{
	"all",
	"analyse",
	"analyze",
	"and",
	"any",
	"array",
	"as",
	"asc",
	"asymmetric",
	"between",
	"bigint",
	"bit",
	"boolean",
	"both",
	"case",
	"cast",
	"char",
	"character",
	"check",
	"coalesce",
	"collate",
	"column",
	"constraint",
	"create",
	"current_catalog",
	"current_date",
	"current_role",
	"current_time",
	"current_timestamp",
	"current_user",
	"dec",
	"decimal",
	"default",
	"deferrable",
	"desc",
	"distinct",
	"do",
	"else",
	"end",
	"except",
	"exists",
	"extract",
	"false",
	"fetch",
	"float",
	"for",
	"foreign",
	"from",
	"grant",
	"greatest",
	"group",
	"grouping",
	"having",
	"in",
	"initially",
	"inout",
	"int",
	"integer",
	"intersect",
	"interval",
	"into",
	"lateral",
	"leading",
	"least",
	"limit",
	"localtime",
	"localtimestamp",
	"national",
	"nchar",
	"none",
	"not",
	"null",
	"nullif",
	"numeric",
	"offset",
	"on",
	"only",
	"or",
	"order",
	"out",
	"overlay",
	"placing",
	"position",
	"precision",
	"primary",
	"real",
	"references",
	"returning",
	"row",
	"select",
	"session_user",
	"setof",
	"smallint",
	"some",
	"substring",
	"symmetric",
	"table",
	"then",
	"time",
	"timestamp",
	"to",
	"trailing",
	"treat",
	"trim",
	"true",
	"union",
	"unique",
	"user",
	"using",
	"values",
	"varchar",
	"variadic",
	"when",
	"where",
	"window",
	"with",
	"xmlattributes",
	"xmlconcat",
	"xmlelement",
	"xmlexists",
	"xmlforest",
	"xmlparse",
	"xmlpi",
	"xmlroot",
	"xmlserialize",
}

// TxDB Interface to wrap methods common to *sql.Tx and *sql.DB
type TxDB interface {
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Database is IGOR
type Database struct {
	connection         TxDB
	db                 TxDB
	rawRows            *sql.Rows
	tables             []string
	joinTables         []string
	models             []DBModel
	logger             *log.Logger
	selectValues       []interface{}
	selectFields       string
	updateCreateValues []interface{}
	updateCreateFields []string
	whereValues        []interface{}
	whereFields        []string
	order              string
	limit              int
	offset             int
	varCount           int
}

// DB returns the current `*sql.DB`
// panics if called during a transaction
func (db *Database) DB() *sql.DB {
	return db.db.(*sql.DB)
}

// Begin initialize a transaction
// panics if begin has been already called
func (db *Database) Begin() *Database {
	// Initialize transaction
	var tx *sql.Tx
	var err error
	if tx, err = db.db.(*sql.DB).Begin(); err != nil {
		db.printLog(err.Error())
		return nil
	}
	// backup db.db into db.connection
	db.connection = db.db.(*sql.DB)
	// replace db.db with the transaction
	db.db = tx
	return db
}

// Commit commits the transaction.
// Panics if the transaction is not started (you have to call Begin before)
func (db *Database) Commit() error {
	err := db.db.(*sql.Tx).Commit()
	// restore connection
	db.db = db.connection
	db.clear()
	return err
}

// Rollback rollbacks the transaction
// Panics if the transaction is not started (you have to call Begin before)
func (db *Database) Rollback() error {
	err := db.db.(*sql.Tx).Rollback()
	// restore connection
	db.db = db.connection
	db.clear()
	return err
}

// DBModel is the interface implemented by every struct that is a relation on the DB
type DBModel interface {
	//TableName returns the associated table name
	TableName() string
}

// clear is called at the end of every query, to clean up the db structure
// preserving the connection and the logger
func (db *Database) clear() {
	db.rawRows = nil
	db.tables = nil
	db.joinTables = nil
	db.models = nil
	db.selectValues = nil
	db.selectFields = ""
	db.updateCreateFields = nil
	db.updateCreateValues = nil
	db.whereValues = nil
	db.whereFields = nil
	db.order = ""
	db.limit = 0
	db.offset = 0
	db.varCount = 1
	// connections contains a backup connection
	// used in trasaction.
	// can be cleaned up
	//db.connection = nil
}

// printLog uses db.log to update log
func (db *Database) printLog(v interface{}) {
	if db.logger != nil {
		db.logger.Print(v)
	}
}

// panicLog uses db.log to update log and than it panics
// if db.log is nil, printLog panic using the panic method
func (db *Database) panicLog(v interface{}) {
	if db.logger != nil {
		db.logger.Panic(v)
	} else {
		panic(v)
	}
}

// Connect opens the connection to PostgreSQL using connectionString
func Connect(connectionString string) (*Database, error) {
	var e error
	db = new(Database)
	if db.db, e = sql.Open("postgres", connectionString); e != nil {
		return nil, e
	}
	db.clear()
	return db, nil
}

// Log sets the query logger
func (db *Database) Log(logger *log.Logger) *Database {
	db.logger = logger
	return db
}

// handleIdentifier handle db idetifiers that are reserved
// It puts reserved keywords used as column/table name betweeen double quotes and
// rename clause into a valid database identifier, following the coventions
func handleIdentifier(clause string) string {
	lowerClause := strings.ToLower(clause)
	i := sort.SearchStrings(reservedKeywords, lowerClause)
	if i < len(reservedKeywords) && reservedKeywords[i] == lowerClause {
		return "\"" + lowerClause + "\""
	}
	return namingConvention(clause)
}

// Model sets the table name for the current query
func (db *Database) Model(model DBModel) *Database {
	db.tables = append(db.tables, handleIdentifier(model.TableName()))
	db.models = append(db.models, model)
	return db
}

// Joins append the join string to the current model
func (db *Database) Joins(joins string) *Database {
	db.joinTables = append(db.joinTables, joins)
	// we can't infer model from the join string (can contain everything)
	return db
}

// Table append the tablename strong to FROM. It has the same behavior of Model, but
// passing the tablename directly as a string
func (db *Database) Table(table string) *Database {
	return db.Joins(table)
}

// Select sets the fields to retrieve. Appends fields to SELECT
func (db *Database) Select(fields string, args ...interface{}) *Database {
	db.selectFields = db.replaceMarks(fields)
	db.selectValues = args
	return db
}

// commonCreateUpdate executs common operation in preparation of create / update statements
// because the logic is the same.
// builder is the function that build the UPDATE or CREATE query
func (db *Database) commonCreateUpdate(value DBModel, builder func() string) error {
	defer db.clear()
	// if Model has been called, skip table name inference procedure
	if len(db.tables) == 0 {
		db.tables = append(db.tables, handleIdentifier(value.TableName()))
		db.models = append(db.models, value)
	}

	in := getStruct(value)
	// getFields handle anonymous nested fields
	fields := getFields(value)

	for _, structField := range fields {
		field := in.FieldByName(structField.Name)
		if value := fieldValue(field, structField); value != nil {
			db.updateCreateFields = append(db.updateCreateFields, getColumnName(structField))
			db.updateCreateValues = append(db.updateCreateValues, value)
		}
	}

	// Compile query
	var stmt *sql.Stmt
	var err error
	if stmt, err = db.db.Prepare(builder()); err != nil {
		return err
	}

	// Pass query parameters and executes the query
	// set db.rawRows to query results (of returning) in order to make it possible to scan rows into result
	if db.rawRows, err = stmt.Query(append(db.updateCreateValues, db.whereValues...)...); err != nil {
		return err
	}
	if err = db.Scan(value); err != nil {
		panic(err)
	}
	return nil
}

// Delete executes DELETE FROM value.TableName where .Where()
// Calling .Where is mandatory. You can pass a nil pointer to value if you just setted
// the tablename with Model.
func (db *Database) Delete(value DBModel) error {
	defer db.clear()
	// if Model has been called, skip table name inference procedure
	if len(db.tables) == 0 {
		db.tables = append(db.tables, handleIdentifier(value.TableName()))
		db.models = append(db.models, value)
	}

	// If where is empty, try to infer a primary key by value
	// otherwise buildDelete panics (where is mandatory)
	db.Where(value)

	// Compile query
	var stmt *sql.Stmt
	var err error
	if stmt, err = db.db.Prepare(db.buildDelete()); err != nil {
		return err
	}

	// Pass query parameters and executes the query
	if _, err = stmt.Exec(db.whereValues...); err != nil {
		return err
	}

	// Clear fields of value after delete, because the object no more exists
	value = reflect.Zero(reflect.ValueOf(value).Type()).Interface().(DBModel)

	return nil
}

// Updates looks for non blank fields in value, extract its value and generate the
// UPDATE value.TableName() SET <field_name> = <value> query part.
// It handles default values when the field is empty.
func (db *Database) Updates(value DBModel) error {
	// Build where condition for update
	err := db.Where(value).commonCreateUpdate(value, db.buildUpdate)
	return err
}

// Create creates a new row into the Database, of type value and with its fields
func (db *Database) Create(value DBModel) error {
	err := db.commonCreateUpdate(value, db.buildCreate)
	return err
}

// Pluck fills the slice with the query result.
// *Executes the query* (calls Scan internally).
// Panics if slice is not a slice or the query is not well formulated
func (db *Database) Pluck(column string, slice interface{}) error {
	dest := reflect.Indirect(reflect.ValueOf(slice))
	if dest.Kind() != reflect.Slice {
		db.panicLog(fmt.Sprintf("slice should be a slice, not %s\n", dest.Kind()))
	}

	db = db.Select(column)
	return db.Scan(slice)
}

// Count sets the query result to be count(*) and scan the result into value.
// *It executes the query* (calls Scan internally).
// It panics if the query is not well formulated.
func (db *Database) Count(value *uint8) error {
	db = db.Select("count(*)")
	return db.Scan(value)
}

// First Scans the result of the selection query of type model using the specified id
// Panics if key is not compatible with the primary key filed type or if the query formulation fails
func (db *Database) First(dest DBModel, key interface{}) error {
	modelKey, _ := primaryKey(dest)
	in := reflect.Indirect(reflect.ValueOf(dest))
	in.FieldByName(modelKey).Set(reflect.Indirect(reflect.ValueOf(key)))
	return db.Model(dest).Where(in.Interface()).Scan(dest)
}

// Scan build the SELECT query and scans the query result query into dest.
// Panics if scan fails or the query fail
func (db *Database) Scan(dest ...interface{}) error {
	ld := len(dest)
	if ld == 0 {
		return errors.New("Required at least one parameter to Scan method")
	}

	var err error
	var rows *sql.Rows
	destIndirect := reflect.Indirect(reflect.ValueOf(dest[0]))
	if db.rawRows == nil {
		// Compile query
		var stmt *sql.Stmt
		// If the destionation is a struct (or a slice of struct)
		// select should select only exported sql fields in the order declared in the struct
		// This thing should go only if the user does not selected with .Select the fields to export
		// Thus only if db.selectFields == ""
		if db.selectFields == "" {
			if ld == 1 { // if is a struct or a slice of struct
				switch destIndirect.Kind() {
				case reflect.Struct:
					db.selectFields = strings.Join(getSQLFields(destIndirect.Interface().(DBModel)), ",")
				case reflect.Slice:
					// hanlde slice of structs and slice of pointers to struct
					sliceType := destIndirect.Type().Elem()
					if sliceType == reflect.Ptr {
						fmt.Println("fucking pointer")
						// TODO
					}
					//sliceType := reflect.Indirect(reflect.New(destIndirect.Type().Elem())).Type().Elem()
					if sliceType.Kind() == reflect.Struct {
						db.selectFields = strings.Join(getSQLFields(reflect.Indirect(reflect.New(sliceType)).Interface().(DBModel)), ",")
					} else {
						panic(sliceType)
					}
				}
			}
		}

		if stmt, err = db.db.Prepare(db.buildSelect()); err != nil {
			return err
		}

		// Pass query parameters and execute it
		if rows, err = stmt.Query(append(db.selectValues, db.whereValues...)...); err != nil {
			return err
		}
	} else {
		rows = db.rawRows
	}

	defer func() {
		rows.Close()
		db.clear()
	}()

	if ld == 1 {
		// if is a slice, find first element to decide how to use scan
		// oterwhise use destIndirect
		var defaultElem reflect.Value
		var slicePtr reflect.Value
		switch destIndirect.Kind() {
		// slice
		case reflect.Slice:
			// create a new element, because slice usually is empty. Thus we have to dinamically create it
			defaultElem = reflect.Indirect(reflect.New(destIndirect.Type().Elem()))
			// Create a pointer to a slice value and set it to the slice
			realSlice := reflect.ValueOf(dest[0])
			slicePtr = reflect.New(realSlice.Type())
		default:
			defaultElem = destIndirect
		}

		// if defaultElem is a struct, extract its fields, pass it to scan (extracts the address)
		var interfaces []interface{}
		if defaultElem.Kind() == reflect.Struct {
			fields := getFields(defaultElem.Interface())
			for _, field := range fields {
				interfaces = append(interfaces, reflect.Indirect(defaultElem.FieldByName(field.Name)).Addr().Interface())
			}
		} else {
			// else convert defaultElem into interfaces, use the address
			interfaces = append(interfaces, defaultElem.Addr().Interface())
		}

		for rows.Next() {
			// defaultElem fields are filled by Scan (scan result into fields as variadic arguments)
			if err = rows.Scan(interfaces...); err != nil {
				return err
			}
			// append result to dest (if the destination is a slice)
			if slicePtr.IsValid() {
				destIndirect.Set(reflect.Append(destIndirect, reflect.Indirect(defaultElem)))
				//x := reflect.Indirect(slicePtr.Elem())
				//x.Set(reflect.Append(x, reflect.ValueOf(defaultElem.Interface())))
				//defaultElem = reflect.Zero(destIndirect.Type().Elem())
			}
		}
	} else {
		// Scan(field1, field2, ...)
		if err = rows.Scan(dest); err != nil {
			return err
		}
	}
	return nil
}

// Raw executes a raw query, replacing placeholders (?) with the one supported by PostgreSQL
// Prepare the statement only. Call Scan to execute itOA
func (db *Database) Raw(query string, args ...interface{}) *Database {
	// Replace ? with $n
	query = db.replaceMarks(query)
	// Append args content to current values
	db.whereValues = append(db.whereValues, args...)

	// Compile query
	var stmt *sql.Stmt
	var err error
	if stmt, err = db.db.Prepare(query); err != nil {
		db.panicLog(err.Error())
	}

	// Pass query parameters and executes the query
	if db.rawRows, err = stmt.Query(db.whereValues...); err != nil {
		db.panicLog(err.Error())
	}
	return db
}

// replaceMarks replace question marks (?) with the PostgreSQL variable identifier
// using the right (incremental) value
func (db *Database) replaceMarks(in string) string {
	var buffer bytes.Buffer
	for _, runeValue := range in {
		if runeValue == '?' {
			buffer.WriteString("$")
			buffer.WriteString(strconv.Itoa(db.varCount))
			db.varCount++
		} else {
			buffer.WriteRune(runeValue)
		}
	}
	return buffer.String()
}

// Where builds the WHERE clause. If a primary key is present in the struct
// only that field is used. Otherwise, every non empty field is ANDed
// s can be a struct, in that case args are ignored
// or it can be a string, in that case args are the query parameters. Use ? placeholder
// If a where condition can't be generated it panics
func (db *Database) Where(s interface{}, args ...interface{}) *Database {
	if reflect.TypeOf(s).Kind() == reflect.String {
		whereClause := reflect.ValueOf(s).String()
		// replace question marks with $n
		// handle cases like .Where("a = ? and b in (?)", 1, []int{1,2,4,6})
		// this must become: a = $1 and b in ($2, $3, $4, $5)
		var slicePos []int

		// since I'm looping through args, I'll build the whereFileds with expanded slices if present
		var whereArgsExtended []interface{}
		for i := 0; i < len(args); i++ {
			if reflect.TypeOf(args[i]).Kind() == reflect.Slice {
				slicePos = append(slicePos, i)
				slice := reflect.Indirect(reflect.ValueOf(args[i]))
				for k := 0; k < slice.Len(); k++ {
					whereArgsExtended = append(whereArgsExtended, reflect.Indirect(slice.Index(k)).Interface())
				}
			} else {
				whereArgsExtended = append(whereArgsExtended, args[i])
			}
		}

		if len(slicePos) > 0 {
			var buffer bytes.Buffer
			// build new where clause, using old where clause until we don't reach the ? associated with the
			// slice. Then replace that ? with len(slice) question marks.
			markCount := 0
			slicePosLen := len(slicePos)
			for _, c := range whereClause {
				if c == '?' {
					s := sort.SearchInts(slicePos, markCount)
					// if found a ? associated with a slice
					if s < slicePosLen && slicePos[s] == markCount {
						sliceLen := reflect.Indirect(reflect.ValueOf(args[markCount])).Len()
						for i := 0; i < sliceLen; i++ {
							buffer.WriteRune('?')
							if i != sliceLen-1 {
								buffer.WriteRune(',')
							}
						}
					} else {
						// if the ? is not associated with a ?, write it as is
						buffer.WriteRune(c)
					}
					markCount++
				} else {
					buffer.WriteRune(c)
				}
			}
			// build the new where clause and pass it to replaceMarks
			fmt.Println(buffer.String())
			db.whereFields = append(db.whereFields, db.replaceMarks(buffer.String()))
			db.whereValues = append(db.whereValues, whereArgsExtended...)
		} else {
			db.whereFields = append(db.whereFields, db.replaceMarks(whereClause))
			db.whereValues = append(db.whereValues, args...)
		}
	} else {
		// must be a struct
		in := getStruct(s)
		key, value := primaryKey(s)

		// if a model has not been setted, set the model as s.TableName()
		if len(db.tables) == 0 {
			db.Model(s.(DBModel))
		}

		escapedTableName := handleIdentifier(s.(DBModel).TableName())

		if key != "" && !isBlank(reflect.ValueOf(value)) {
			db.whereFields = append(db.whereFields, escapedTableName+"."+handleIdentifier(key))
			db.whereValues = append(db.whereValues, value)
		} else {
			for i := 0; i < in.NumField(); i++ {
				fieldValue := in.Field(i)
				fieldType := in.Type().Field(i)
				if !isBlank(fieldValue) {
					db.whereFields = append(db.whereFields, escapedTableName+"."+getColumnName(fieldType))
					db.whereValues = append(db.whereValues, fieldValue.Interface())
				}
			}
		}
	}
	return db
}

// primaryKey returns the pair (key, value) representing the defined primary key field
// (key) and its value (value), when the `sql` struct tag field is defined and is value is not blank
// returns empty key if a key is not present and thus en empty value
// oterwise returns the key and the value (if present, i.e. is not blank)
// returned Key is the Name of the field. Not following the sql conventions but the go convention.
// If you need to change this value (and you usually do), parse key with handleIdentifier
func primaryKey(s interface{}) (key string, value interface{}) {
	val := reflect.Indirect(reflect.ValueOf(s))
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := val.Type().Field(i)
		// Handle embedded anonymous struct only
		switch fieldValue.Kind() {
		case reflect.Struct:
			if fieldType.Anonymous {
				key, value = primaryKey(fieldValue.Interface())
			}
		default:
			tag := strings.ToLower(fieldType.Tag.Get("gorm"))
			tagValue := strings.Split(tag, ",")
			if utils.InSlice("primary_key", tagValue) {
				key = fieldType.Name
				value = fieldValue.Interface()
				return
			}
		}
	}
	return
}

// namingConvention returns the coversion of input name to a
// valid db entity that follows the convention
func namingConvention(name string) string {
	// first char is always upper case
	var ucActual = true
	var buffer bytes.Buffer
	buffer.WriteByte(name[0])
	for i := 1; i < len(name); i++ {
		prevChar := rune(name[i-1])
		actualChar := rune(name[i])
		ucActual = unicode.IsUpper(actualChar)
		if unicode.IsLower(prevChar) && ucActual {
			buffer.WriteByte('_')
		}
		buffer.WriteRune(actualChar)
	}

	return strings.ToLower(buffer.String())
}

// getColumnName returns the column name of the specified field of the struct
// it's the name of the field if the field has not a `gorm:column` value spcified
// the field is a valid sql value (thus in case, the name is escaped using handleIdentifier)
func getColumnName(field reflect.StructField) (fieldName string) {
	ts := parseTagSetting(field.Tag.Get("gorm"))
	if ts["column"] != "" {
		fieldName = ts["column"]
	} else {
		fieldName = handleIdentifier(field.Name)
	}
	return
}

// getSQLFields returns sql-compatible fields that the select query should return
// skips sql:"-".
func getSQLFields(s DBModel) (ret []string) {
	fields := getFields(s)
	table := handleIdentifier(s.TableName())
	for _, field := range fields {
		ret = append(ret, table+"."+getColumnName(field))
	}
	return
}

// getFields returns a slice of reflect.StructField that represents the exported struct Fields in s
// that are not excluded in sql generation
func getFields(s interface{}) (ret []reflect.StructField) {
	val := reflect.Indirect(reflect.ValueOf(s))
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := val.Type().Field(i)
		// if it's exported
		if fieldType.PkgPath == "" {
			// Handle embedded anoymous struct
			switch fieldValue.Kind() {
			case reflect.Struct:
				// if it's anonymous, embed its fields in the query
				if fieldType.Anonymous {
					ret = append(ret, getFields(fieldValue.Interface())...)
				} else { // use its name only (to work with structs like time.Time)
					tag := strings.ToLower(fieldType.Tag.Get("sql"))
					tagValue := strings.Split(tag, ",")
					if !utils.InSlice("-", tagValue) {
						ret = append(ret, fieldType)
					}
				}
			default:
				tag := strings.ToLower(fieldType.Tag.Get("sql"))
				tagValue := strings.Split(tag, ",")
				if !utils.InSlice("-", tagValue) {
					ret = append(ret, fieldType)
				}
			}
		}
	}

	return
}

// isBlank returns true if value is empty
func isBlank(value reflect.Value) bool {
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// getStruct panics if s is not a struct. Returns the reflect.Indirect(reflect.ValueOf(s))
func getStruct(s interface{}) reflect.Value {
	in := reflect.Indirect(reflect.ValueOf(s))
	if in.Kind() != reflect.Struct {
		panic(fmt.Sprintf("s must be a sturct, not %s\n", in.Kind()))
	}
	return in
}

// buildSelect returns the generated SQL. Panics if it can't generate a query
func (db *Database) buildSelect() string {
	if len(db.tables) == 0 {
		db.panicLog("Please set a table with Model [ + Joins ]")
	}

	var query bytes.Buffer
	// Select
	var fields string
	query.WriteString("SELECT ")
	if len(db.selectFields) > 0 {
		fields = db.selectFields
	}

	query.WriteString(fields)
	query.WriteString(" FROM ")
	// Model && Join

	query.WriteString(strings.Join(db.tables, ","))
	query.WriteString(" ")

	// Join (optional)
	if len(db.joinTables) > 0 {
		query.WriteString(strings.Join(db.joinTables, " "))
		query.WriteString(" ")
	}

	// Where (optional)
	if len(db.whereFields) > 0 {
		query.WriteString(db.buildWhere())
	}

	// Order (optional)
	if db.order != "" {
		query.WriteString(" ORDER BY ")
		query.WriteString(db.order)
	}

	// Limit (optional)
	if db.limit != 0 {
		query.WriteString(" LIMIT ")
		query.WriteString(strconv.Itoa(db.limit))
	}

	// Offset (optional)
	if db.offset != 0 {
		query.WriteString(" OFFSET ")
		query.WriteString(strconv.Itoa(db.limit))
	}

	query.WriteString(";")
	qs := query.String()
	db.printLog(qs)
	return qs
}

// Limit sets the LIMIT value to the query
func (db *Database) Limit(limit int) *Database {
	db.limit = limit
	return db
}

// Offset sets the OFFSET value to the query
func (db *Database) Offset(offset int) *Database {
	db.offset = offset
	return db
}

// Order sets the ORDER BY value to the query
func (db *Database) Order(value string) *Database {
	db.order = handleIdentifier(value)
	return db
}

// buildUpdate returns the generated SQL for the UPDATE statement. Panics if it can't generate a query
func (db *Database) buildUpdate() string {
	var query bytes.Buffer
	query.WriteString("UPDATE ")

	// Model only
	if len(db.tables) != 1 {
		db.panicLog("Please set a table with Model to Update")
	}
	query.WriteString(db.tables[0])
	query.WriteString(" SET ")

	updateSize := len(db.updateCreateFields)
	if updateSize == 0 {
		db.panicLog("Unable to detect fields to update")
	}

	for j, field := range db.updateCreateFields {
		query.WriteString(handleIdentifier(field))
		query.WriteString(" = $")
		query.WriteString(strconv.Itoa(db.varCount))
		db.varCount++
		if j != updateSize-1 {
			query.WriteString(", ")
		}
	}

	// Where (optional)
	if len(db.whereFields) > 0 {
		query.WriteString(db.buildWhere())
	}

	query.WriteString(db.buildReturning())

	qs := query.String()
	db.printLog(qs)
	return qs
}

// buildCreate returns the generated SQL for the CREATE statement. Panics if it can't generate a query
func (db *Database) buildCreate() string {
	var query bytes.Buffer
	query.WriteString("INSERT INTO ")

	// Model only
	if len(db.tables) != 1 {
		db.panicLog(fmt.Sprintf("Unable to infer table name for Create. Number of tables: %d", len(db.tables)))
	}
	// Table (
	query.WriteString(db.tables[0])
	query.WriteString("(")

	// field1,file2,...
	createSize := len(db.updateCreateFields)
	if createSize == 0 {
		db.panicLog("Unable to detect fields for Create")
	}

	query.WriteString(strings.Join(db.updateCreateFields, ","))
	// values($1,..$n) place holders
	query.WriteString(") VALUES (")
	for i := 0; i < createSize; i++ {
		query.WriteString("$")
		query.WriteString(strconv.Itoa(db.varCount))
		db.varCount++
		if i != createSize-1 {
			query.WriteString(",")
		}
	}

	query.WriteString(") ")
	query.WriteString(db.buildReturning())
	query.WriteString(";")

	qs := query.String()
	db.printLog(qs)
	return qs
}

// buildReturning returns the RETURNING part of the query
// it explicits every fields in the current model.
// In that way we're able to easily Scan the results
func (db *Database) buildReturning() string {
	var query bytes.Buffer
	query.WriteString(" RETURNING ")
	query.WriteString(strings.Join(getSQLFields(db.models[0]), ","))
	return query.String()
}

// buildDelete returns the generated SQL for the DELETE statement. Panics if it can't geerate a query
func (db *Database) buildDelete() string {
	var query bytes.Buffer
	query.WriteString("DELETE FROM ")

	// Model only
	if len(db.tables) != 1 {
		db.panicLog("Unable to infer table name for Delete. Use Delete(model) or Model(model)")
	}
	query.WriteString(db.tables[0])

	// Where (mandatory)
	if len(db.whereFields) == 0 {
		db.panicLog("Where statement is mandatory in Delete")
	}

	query.WriteString(db.buildWhere())
	query.WriteString(";")

	qs := query.String()
	db.printLog(qs)
	return qs
}

// buildWhere returns the generated SQL for the WHERE clause. Panics if the Where method hasn't been called
// The generated query uses the PostgreSQL placeholders for query parameters in compiled queries
func (db *Database) buildWhere() string {
	var query bytes.Buffer
	whereSize := len(db.whereFields)
	if whereSize == 0 {
		db.panicLog("Please add a Where condition with .Where")
	}

	query.WriteString(" WHERE ")

	for j, clause := range db.whereFields {
		if strings.Contains(clause, "$") {
			query.WriteString(clause)
		} else {
			query.WriteString(handleIdentifier(clause))
			query.WriteString(" = $")
			query.WriteString(strconv.Itoa(db.varCount))
			db.varCount++
		}
		if j != whereSize-1 {
			query.WriteString(" AND ")
		}
	}
	query.WriteString(" ")
	return query.String()
}

// escapeIfNeeded escapes the value if needed in query generation
func escapeIfNeeded(value string) string {
	trimmed := strings.TrimSpace(value)
	// default:'string value' OR
	//sql expression, like: default:"(now() at timezone 'utc') or now() or user_defined_function(parameters.. )
	if (strings.HasPrefix(trimmed, "'") && strings.HasSuffix(trimmed, "'")) ||
		strings.HasSuffix(trimmed, ")") {
		return trimmed
	}

	lowered := strings.ToLower(trimmed)
	// null and other sql reserved keyworks (used a default values) can't be placed between single quotes
	if lowered == "null" || strings.HasPrefix(lowered, "current_") {
		return lowered
	}
	// default:'something' like:default:'false' should be between single quotes
	return "'" + trimmed + "'"
}

// fieldValue returns an interface{} that's the value of fieldVal if fieldVal is not blank
// if fieldVal is blank and the field has a default value, return the defalt value
// oherwise returns nil
func fieldValue(fieldVal reflect.Value, structField reflect.StructField) (ret interface{}) {
	if isBlank(fieldVal) {
		defaultValue := strings.TrimSpace(parseTagSetting(structField.Tag.Get("sql"))["default"])
		if defaultValue != "" {
			switch fieldVal.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if numericValue, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
					if numericValue != fieldVal.Int() {
						return fieldVal.Int()
					}
					return numericValue
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if numericValue, err := strconv.ParseUint(defaultValue, 10, 64); err == nil {
					if numericValue != fieldVal.Uint() {
						return fieldVal.Int()
					}
					return numericValue
				}
			case reflect.Float32, reflect.Float64:
				if floatValue, err := strconv.ParseFloat(defaultValue, 64); err == nil {
					if floatValue != fieldVal.Float() {
						return fieldVal.Float()
					}
					return floatValue
				}
			case reflect.Bool:
				if boolValue, err := strconv.ParseBool(defaultValue); err == nil {
					if boolValue != fieldVal.Bool() {
						return fieldVal.Bool()
					}
					return boolValue
				}
			case reflect.String:
				return fieldVal.String()
			default:
				return nil
			}
		}
		return nil
	}
	return fieldVal.Interface()
}

// parseTagSetting, imported from jinzhu/gorm
func parseTagSetting(str string) map[string]string {
	tags := strings.Split(str, ";")
	setting := map[string]string{}
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToLower(v[0]))
		if len(v) >= 2 {
			setting[k] = strings.Join(v[1:], ":")
		} else {
			setting[k] = k
		}
	}
	return setting
}
