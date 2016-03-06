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
	// Blank import required to get PostgreSQL working
	_ "github.com/lib/pq"
	"github.com/nerdzeu/nerdz-api/utils"
)

// Reserved keyworkds that shouldn't be used as column name or other identifiers
// result of: select '"' || word::text || '",' from pg_get_keywords() where catcode = 'R' OR catdesc like '%cannot%';
// We will check if the current identifier is in that list. If it is, it will be placed between quotes
var reserved_keywords = []string{
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

// TxDB Interface to wrap methods common to Tx and Database
type TxDB interface {
	Prepare(query string) (*sql.Stmt, error)
}

// Database is IGOR
type Database struct {
	db                 *sql.DB
	rawRows            *sql.Rows
	tables             []string
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

// Tx is a transaction
type Tx struct {
	Database
	tx *sql.Tx
}

// Begin initialize a transaction
func (db *Database) Begin() (tx *Tx) {
	// copy database connection
	tx.db = db.db
	var err error
	// Initialize transaction
	if tx.tx, err = tx.db.Begin(); err != nil {
		db.printLog(err.Error())
		return nil
	}
	return
}

// Commit commits the transaction
func (tx *Tx) Commit() error {
	//TODO
	return nil
}

// Rollback rollbacks the transaction
func (tx *Tx) Rollback() error {
	//TODO
	return nil
}

// DBModel is the interface implemented by every struct that is a relation on the DB
type DBModel interface {
	//TableName returns the associated table name
	TableName() string
}

// clear is called at the end of every query, to clean up the db structure
// preserving the connection and the logger
func (db *Database) clear() {
	db.tables = nil
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

// Ping pings the Database to check Database connection
func (db *Database) Ping() error {
	return db.db.Ping()
}

// Log sets the query logger
func (db *Database) Log(logger *log.Logger) *Database {
	db.logger = logger
	return db
}

// handleReserved handle db idetifiers that are reserved
// It puts reserved keywords used as column/table name between double quotes
func handleReserved(clause string) string {
	lowerClause := strings.ToLower(clause)
	i := sort.SearchStrings(reserved_keywords, lowerClause)
	if i < len(reserved_keywords) && reserved_keywords[i] == lowerClause {
		lowerClause = "\"" + lowerClause + "\""
	}
	return lowerClause

}

// Model sets the table name for the current query
func (db *Database) Model(model DBModel) *Database {
	db.tables = append(db.tables, handleReserved(model.TableName()))
	db.models = append(db.models, model)
	return db
}

// Joins append the join string to the current model
func (db *Database) Joins(joins string) *Database {
	db.tables = append(db.tables, joins)
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
	db.selectFields = fields
	db.selectValues = args
	return db
}

// commonCreateUpdate executs common operation in preparation of create / update statements
// because the logic is the same.
// query is the generated query for update / create
func (db *Database) commonCreateUpdate(query string, value DBModel) error {
	// if Model has been called, skip table name inference procedure
	if len(db.tables) == 0 {
		db.tables = append(db.tables, handleReserved(value.TableName()))
		db.models = append(db.models, value)
	}
	in := getStruct(value)
	Type := reflect.TypeOf(in)
	for i := 0; i < in.NumField(); i++ {
		structField := Type.Field(i)
		// fieldValue
		stringValue := fieldValue(reflect.ValueOf(in.Field(i).Elem().Interface()), structField)
		if stringValue != "" {
			db.updateCreateFields = append(db.updateCreateFields, structField.Name)
			db.updateCreateValues = append(db.updateCreateValues, stringValue)
		}
	}

	// Compile query
	var stmt *sql.Stmt
	var err error
	if stmt, err = db.db.Prepare(query); err != nil {
		return err
	}

	// Pass query parameters and executes the query
	if _, err = stmt.Exec(db.updateCreateValues, db.whereValues); err != nil {
		return err
	}
	db.clear()
	return nil
}

// Delete executes DELETE FROM value.TableName where .Where()
// Calling .Where is mandatory. You can pass a nil pointer to value if you just setted
// the tablename with Model.
func (db *Database) Delete(value DBModel) error {
	// if Model has been called, skip table name inference procedure
	if len(db.tables) == 0 {
		db.tables = append(db.tables, handleReserved(value.TableName()))
		db.models = append(db.models, value)
	}

	// Compile query
	var stmt *sql.Stmt
	var err error
	if stmt, err = db.db.Prepare(db.buildDelete()); err != nil {
		return err
	}

	// Pass query parameters and executes the query
	if _, err = stmt.Exec(db.whereValues); err != nil {
		return err
	}
	db.clear()
	// Clear fields of value after delete, because the object no more exists
	value = reflect.New(reflect.TypeOf(value)).Interface().(DBModel)

	return nil
}

// Updates looks for non blank fields in value, extract its value and generate the
// UPDATE value.TableName() SET <field_name> = <value> query part.
// It handles default values when the field is empty.
func (db *Database) Updates(value DBModel) error {
	return db.commonCreateUpdate(db.buildUpdate(), value)
}

// Create creates a new row into the Database, of type value and with its fields
func (db *Database) Create(value DBModel) error {
	return db.commonCreateUpdate(db.buildCreate(), value)
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
					fields := getSelectFields(destIndirect.Interface())
					db.selectFields = strings.Join(fields, ",")
				case reflect.Slice:
					if destIndirect.Index(0).Kind() == reflect.Struct {
						fields := getSelectFields(reflect.Indirect(destIndirect.Index(0)).Interface())
						db.selectFields = strings.Join(fields, ",")
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
		switch destIndirect.Kind() {
		// slice
		case reflect.Slice:
			defaultElem = reflect.Indirect(destIndirect.Index(0))
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
			if err = rows.Scan(interfaces...); err != nil {
				return err
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
	db.clear()
	// Replace ? with $n
	db.Where(query, args)

	// Compile query
	var stmt *sql.Stmt
	var err error
	if stmt, err = db.db.Prepare(strings.Join(db.whereFields, " ")); err != nil {
		db.panicLog(err.Error())
	}

	// Pass query parameters and executes the query
	if db.rawRows, err = stmt.Query(db.whereValues); err != nil {
		db.panicLog(err.Error())
	}
	return db
}

// Where builds the WHERE clause. If a primary key is present in the struct
// only that field is used. Otherwise, every non empty field is ANDed
// s can be a struct, in that case args are ignored
// or it can be a string, in that case args are the query parameters. Use ? placeholder
// If a where condition can't be generated it panics
func (db *Database) Where(s interface{}, args ...interface{}) *Database {
	if reflect.TypeOf(s).Kind() == reflect.String {
		whereClause := reflect.ValueOf(s).String()
		var buffer bytes.Buffer
		for _, runeValue := range whereClause {
			if runeValue == '?' {
				buffer.WriteString("$")
				buffer.WriteString(strconv.Itoa(db.varCount))
				db.varCount++
			} else {
				buffer.WriteRune(runeValue)
			}
		}
		db.whereFields = append(db.whereFields, buffer.String())
		db.whereValues = append(db.whereValues, args)
		whereSize := len(db.whereFields)
		if whereSize == 0 {
			db.panicLog(fmt.Sprintf("can't build where conidition with string: %s", whereClause))
		}
	} else {
		in := getStruct(s)
		key, value := primaryKey(s)

		if key != "" && !isBlank(reflect.ValueOf(value)) {
			db.whereFields = append(db.whereFields, key)
			db.whereValues = append(db.whereValues, value)
		} else {
			Type := reflect.TypeOf(in)
			for i := 0; i < in.NumField(); i++ {
				field := in.Field(i)
				value := field.Elem().Interface()
				if !isBlank(reflect.ValueOf(value)) {
					db.whereFields = append(db.whereFields, Type.Field(i).Name)
					db.whereValues = append(db.whereValues, value)
				}
			}
		}
		whereSize := len(db.whereFields)
		if whereSize == 0 {
			db.panicLog(fmt.Sprintf("can't build where condition with struct: %v\n", in))
		}
	}

	return db
}

// primaryKey returns the pair (key, value) representing the defined primary key field
// (key) and its value (value), when the `sql` struct tag field is defined and is value is not blank
// returns empty key if a key is not present and thus en empty value
// oterwise returns the key and the value (if present, i.e. is not blank)
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

// getSelectFields returns sql-compatible fields that the select query should return
// skips sql:"-".
func getSelectFields(s interface{}) (ret []string) {
	fields := getFields(s)
	for _, field := range fields {
		ret = append(ret, handleReserved(field.Name))
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
	var query bytes.Buffer
	// Select
	fields := "*"
	query.WriteString("SELECT ")
	if len(db.selectFields) > 0 {
		fields = db.selectFields
	}

	query.WriteString(fields)
	query.WriteString(" FROM ")
	// Model && Join
	if len(db.tables) == 0 {
		db.panicLog("Please set a table with Model [ + Joins ]")
	}
	query.WriteString(strings.Join(db.tables, ","))

	// Where (optional)
	if len(db.whereFields) > 0 {
		query.WriteString(db.buildWhere())
	}

	// Order (optional)
	if db.order != "" {
		query.WriteString(" ")
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
	db.order = value
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
		query.WriteString(handleReserved(field))
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

	query.WriteString(" RETURNING ")
	query.WriteString(db.tables[0])
	query.WriteString(".")
	key, _ := primaryKey(db.models[0])
	query.WriteString(key)
	query.WriteString(";")

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
		db.panicLog("Unable to infer table name for Create")
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
		if i != createSize {
			query.WriteString(",")
		}
	}
	query.WriteString(");")

	qs := query.String()
	db.printLog(qs)
	return qs
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

	query.WriteString(" WHERE (")

	for j, clause := range db.whereFields {
		if strings.Contains(clause, "$") {
			query.WriteString(clause)
		} else {
			query.WriteString(handleReserved(clause))
			query.WriteString(" = $")
			query.WriteString(strconv.Itoa(db.varCount))
			db.varCount++
		}
		if j != whereSize-1 {
			query.WriteString(" AND ")
		}
	}
	query.WriteString(") ")
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

// fieldValue returns the default value of the structField if fieldVal is blank
// Otherwise fieldVal is convertend to string and returned
func fieldValue(fieldVal reflect.Value, structField reflect.StructField) string {
	if isBlank(fieldVal) {
		defaultValue := strings.TrimSpace(parseTagSetting(structField.Tag.Get("sql"))["default"])
		switch fieldVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if numericValue, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
				if numericValue != fieldVal.Int() {
					return escapeIfNeeded(fmt.Sprintf("%d", fieldVal.Int()))
				} else {
					return escapeIfNeeded(defaultValue)
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if numericValue, err := strconv.ParseUint(defaultValue, 10, 64); err == nil {
				if numericValue != fieldVal.Uint() {
					return escapeIfNeeded(escapeIfNeeded(fmt.Sprintf("%d", fieldVal.Int())))
				} else {
					return escapeIfNeeded(defaultValue)
				}
			}
		case reflect.Float32, reflect.Float64:
			if floatValue, err := strconv.ParseFloat(defaultValue, 64); err == nil {
				if floatValue != fieldVal.Float() {
					return escapeIfNeeded(fmt.Sprintf("%f", fieldVal.Float()))
				} else {
					return escapeIfNeeded(defaultValue)
				}
			}
		case reflect.Bool:
			if boolValue, err := strconv.ParseBool(defaultValue); err == nil {
				if boolValue != fieldVal.Bool() {
					return escapeIfNeeded(fmt.Sprintf("%t", fieldVal.Bool()))
				} else {
					return escapeIfNeeded(defaultValue)
				}
			}
		default:
			return escapeIfNeeded(defaultValue)
		}
	} else {
		switch fieldVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return escapeIfNeeded(fmt.Sprintf("%d", fieldVal.Int()))
		case reflect.Bool:
			return escapeIfNeeded(fmt.Sprintf("%t", fieldVal.Bool()))
		case reflect.Float32, reflect.Float64:
			return escapeIfNeeded(fmt.Sprintf("%f", fieldVal.Float()))
		case reflect.String:
			return escapeIfNeeded(fieldVal.String())
		default:
			panic(fmt.Sprintf("Unsupported type %s in sql generation", fieldVal.Kind()))
		}
	}
	return ""
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
