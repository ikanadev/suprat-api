package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Color holds the schema definition for the Color entity.
type Color struct {
	ent.Schema
}

// Fields of the Color.
func (Color) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("rgb"),
	}
}

// Edges of the Color.
func (Color) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities", Activity.Type),
		edge.To("measurements", Measurement.Type),
	}
}
