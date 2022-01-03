package schema

import (
	"annie-baby/enum"
	"github.com/iawia002/annie/extractors/types"
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
	enums := enum.VideoEnum.GetEnumsMsgs()
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
			Values(enums...).Default(enum.Pending.Msg),

		field.JSON("streams", map[string]*types.Stream{}).Optional(),

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
