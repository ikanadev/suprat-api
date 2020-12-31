package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Measure holds the schema definition for the Measure entity.
type Measure struct {
	ent.Schema
}

// Fields of the Measure.
func (Measure) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

// Edges of the Measure.
func (Measure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("record", Record.Type).Ref("measures").Unique().Required(),
		edge.From("measurement", Measurement.Type).Ref("measures").Unique().Required(),
	}
}
