package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Measurement holds the schema definition for the Measurement entity.
type Measurement struct {
	ent.Schema
}

// Fields of the Measurement.
func (Measurement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("enabled").Default(true),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Measurement.
func (Measurement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("measurements").Unique().Required(),
		edge.From("color", Color.Type).Ref("measurements").Unique().Required(),
		edge.From("activities", Activity.Type).Ref("measurements"),

		edge.To("measures", Measure.Type),
	}
}
