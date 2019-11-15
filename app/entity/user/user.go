// Code generated by entc, DO NOT EDIT.

package user

import (
	"gostalt/app/entity/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	fields = schema.User{}.Fields()

	// descName is the schema descriptor for name field.
	descName = fields[0].Descriptor()
	// DefaultName holds the default value on creation for the name field.
	DefaultName = descName.Default.(string)
)