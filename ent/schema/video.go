package schema

import (
	"annie-baby/enum"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("link").
			NotEmpty().
			Unique(),

		// 网站类型, 不用enum了
		field.String("type").
			NotEmpty(),

		field.String("title").
			NotEmpty(),

		field.String("output_dir"),

		field.String("proxy"),

		field.Enum("status").
			Values(
				enum.PENDING.String(),
			),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return nil
}
