// Code generated by entc, DO NOT EDIT.

package video

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldOutputDir holds the string denoting the output_dir field in the database.
	FieldOutputDir = "output_dir"
	// FieldProxy holds the string denoting the proxy field in the database.
	FieldProxy = "proxy"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the video in the database.
	Table = "videos"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldLink,
	FieldType,
	FieldTitle,
	FieldOutputDir,
	FieldProxy,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

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
	// LinkValidator is a validator for the "link" field. It is called by the builders before save.
	LinkValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusPENDING Status = "PENDING"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPENDING:
		return nil
	default:
		return fmt.Errorf("video: invalid enum value for status field: %q", s)
	}
}
