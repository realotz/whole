// Code generated by entc, DO NOT EDIT.

package employee

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the employee type in the database.
	Label = "employee"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldIDCard holds the string denoting the id_card field in the database.
	FieldIDCard = "id_card"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldLastIP holds the string denoting the last_ip field in the database.
	FieldLastIP = "last_ip"
	// FieldLastTime holds the string denoting the last_time field in the database.
	FieldLastTime = "last_time"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// Table holds the table name of the employee in the database.
	Table = "employees"
	// RolesTable is the table the holds the roles relation/edge. The primary key declared below.
	RolesTable = "employee_roles"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "roles"
)

// Columns holds all SQL columns for employee fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSex,
	FieldUUID,
	FieldAccount,
	FieldAvatar,
	FieldName,
	FieldEmail,
	FieldMobile,
	FieldIDCard,
	FieldBirthday,
	FieldPassword,
	FieldSalt,
	FieldLastIP,
	FieldLastTime,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"employee_id", "role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// AccountValidator is a validator for the "account" field. It is called by the builders before save.
	AccountValidator func(string) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	SaltValidator func(string) error
)

// Sex defines the type for the "sex" enum field.
type Sex string

// SexUnknown is the default value of the Sex enum.
const DefaultSex = SexUnknown

// Sex values.
const (
	SexUnknown Sex = "unknown"
	SexMan     Sex = "man"
	SexWoman   Sex = "woman"
)

func (s Sex) String() string {
	return string(s)
}

// SexValidator is a validator for the "sex" field enum values. It is called by the builders before save.
func SexValidator(s Sex) error {
	switch s {
	case SexUnknown, SexMan, SexWoman:
		return nil
	default:
		return fmt.Errorf("employee: invalid enum value for sex field: %q", s)
	}
}
