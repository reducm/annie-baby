// Code generated by entc, DO NOT EDIT.

package ent

import (
	"annie-baby/ent/schema"
	"annie-baby/ent/video"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescLink is the schema descriptor for link field.
	videoDescLink := videoFields[0].Descriptor()
	// video.LinkValidator is a validator for the "link" field. It is called by the builders before save.
	video.LinkValidator = videoDescLink.Validators[0].(func(string) error)
	// videoDescType is the schema descriptor for type field.
	videoDescType := videoFields[1].Descriptor()
	// video.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	video.TypeValidator = videoDescType.Validators[0].(func(string) error)
	// videoDescTitle is the schema descriptor for title field.
	videoDescTitle := videoFields[2].Descriptor()
	// video.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	video.TitleValidator = videoDescTitle.Validators[0].(func(string) error)
	// videoDescCreatedAt is the schema descriptor for created_at field.
	videoDescCreatedAt := videoFields[6].Descriptor()
	// video.DefaultCreatedAt holds the default value on creation for the created_at field.
	video.DefaultCreatedAt = videoDescCreatedAt.Default.(func() time.Time)
	// videoDescUpdatedAt is the schema descriptor for updated_at field.
	videoDescUpdatedAt := videoFields[7].Descriptor()
	// video.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	video.DefaultUpdatedAt = videoDescUpdatedAt.Default.(func() time.Time)
	// video.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	video.UpdateDefaultUpdatedAt = videoDescUpdatedAt.UpdateDefault.(func() time.Time)
}
