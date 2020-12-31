package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Icon holds the schema definition for the Icon entity.
type Icon struct {
	ent.Schema
}

// Fields of the Icon.
func (Icon) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Icon.
func (Icon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities", Activity.Type),
	}
}
