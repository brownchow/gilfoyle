// Code generated by entc, DO NOT EDIT.

package media

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the media in the database.
	Table = "media"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusProcessing Status = "processing"
	StatusReady      Status = "ready"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusProcessing, StatusReady:
		return nil
	default:
		return fmt.Errorf("media: invalid enum value for status field: %q", s)
	}
}
