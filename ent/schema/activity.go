package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("activities").Unique().Required(),
		edge.From("color", Color.Type).Ref("activities").Unique().Required(),
		edge.From("icon", Icon.Type).Ref("activities").Unique().Required(),

		edge.To("records", Record.Type),
		edge.To("measurements", Measurement.Type),
	}
}
