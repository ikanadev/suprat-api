// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/suprat-api/ent/activity"
	"github.com/vmkevv/suprat-api/ent/color"
	"github.com/vmkevv/suprat-api/ent/measurement"
)

// ColorCreate is the builder for creating a Color entity.
type ColorCreate struct {
	config
	mutation *ColorMutation
	hooks    []Hook
}

// SetName sets the name field.
func (cc *ColorCreate) SetName(s string) *ColorCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetRgb sets the rgb field.
func (cc *ColorCreate) SetRgb(s string) *ColorCreate {
	cc.mutation.SetRgb(s)
	return cc
}

// AddActivityIDs adds the activities edge to Activity by ids.
func (cc *ColorCreate) AddActivityIDs(ids ...int) *ColorCreate {
	cc.mutation.AddActivityIDs(ids...)
	return cc
}

// AddActivities adds the activities edges to Activity.
func (cc *ColorCreate) AddActivities(a ...*Activity) *ColorCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddActivityIDs(ids...)
}

// AddMeasurementIDs adds the measurements edge to Measurement by ids.
func (cc *ColorCreate) AddMeasurementIDs(ids ...int) *ColorCreate {
	cc.mutation.AddMeasurementIDs(ids...)
	return cc
}

// AddMeasurements adds the measurements edges to Measurement.
func (cc *ColorCreate) AddMeasurements(m ...*Measurement) *ColorCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cc.AddMeasurementIDs(ids...)
}

// Mutation returns the ColorMutation object of the builder.
func (cc *ColorCreate) Mutation() *ColorMutation {
	return cc.mutation
}

// Save creates the Color in the database.
func (cc *ColorCreate) Save(ctx context.Context) (*Color, error) {
	var (
		err  error
		node *Color
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ColorCreate) SaveX(ctx context.Context) *Color {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *ColorCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := cc.mutation.Rgb(); !ok {
		return &ValidationError{Name: "rgb", err: errors.New("ent: missing required field \"rgb\"")}
	}
	return nil
}

func (cc *ColorCreate) sqlSave(ctx context.Context) (*Color, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ColorCreate) createSpec() (*Color, *sqlgraph.CreateSpec) {
	var (
		_node = &Color{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: color.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: color.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: color.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Rgb(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: color.FieldRgb,
		})
		_node.Rgb = value
	}
	if nodes := cc.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   color.ActivitiesTable,
			Columns: []string{color.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: activity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.MeasurementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   color.MeasurementsTable,
			Columns: []string{color.MeasurementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: measurement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ColorCreateBulk is the builder for creating a bulk of Color entities.
type ColorCreateBulk struct {
	config
	builders []*ColorCreate
}

// Save creates the Color entities in the database.
func (ccb *ColorCreateBulk) Save(ctx context.Context) ([]*Color, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Color, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ColorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ccb *ColorCreateBulk) SaveX(ctx context.Context) []*Color {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}