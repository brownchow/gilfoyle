package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// MediaStatusProcessing relates to the initial state of a media
const MediaStatusProcessing = "processing"

// MediaStatusReady relates to the final state of a media
const MediaStatusReady = "ready"

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

func (Media) Config() ent.Config {
	return ent.Config{
		Table: "media",
	}
}

func (Media) Index() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(),
	}
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(func() uuid.UUID {
			return uuid.New()
		}),
		field.String("title").NotEmpty().MinLen(1).MaxLen(255),
		field.Enum("status").Values(MediaStatusProcessing, MediaStatusReady),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("updated_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return nil
}
