// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent/schema"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[3].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
