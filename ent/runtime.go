// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/omecodes/omestore/ent/access"
	"github.com/omecodes/omestore/ent/group"
	"github.com/omecodes/omestore/ent/permission"
	"github.com/omecodes/omestore/ent/schema"
	"github.com/omecodes/omestore/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	accessFields := schema.Access{}.Fields()
	_ = accessFields
	// accessDescCreator is the schema descriptor for creator field.
	accessDescCreator := accessFields[1].Descriptor()
	// access.CreatorValidator is a validator for the "creator" field. It is called by the builders before save.
	access.CreatorValidator = accessDescCreator.Validators[0].(func(string) error)
	// accessDescID is the schema descriptor for id field.
	accessDescID := accessFields[0].Descriptor()
	// access.IDValidator is a validator for the "id" field. It is called by the builders before save.
	access.IDValidator = accessDescID.Validators[0].(func(string) error)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.IDValidator is a validator for the "id" field. It is called by the builders before save.
	group.IDValidator = groupDescID.Validators[0].(func(string) error)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescUser is the schema descriptor for user field.
	permissionDescUser := permissionFields[1].Descriptor()
	// permission.UserValidator is a validator for the "user" field. It is called by the builders before save.
	permission.UserValidator = permissionDescUser.Validators[0].(func(string) error)
	// permissionDescData is the schema descriptor for data field.
	permissionDescData := permissionFields[2].Descriptor()
	// permission.DataValidator is a validator for the "data" field. It is called by the builders before save.
	permission.DataValidator = permissionDescData.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescValidated is the schema descriptor for validated field.
	userDescValidated := userFields[3].Descriptor()
	// user.DefaultValidated holds the default value on creation for the validated field.
	user.DefaultValidated = userDescValidated.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
