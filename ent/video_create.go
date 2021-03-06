// Code generated by entc, DO NOT EDIT.

package ent

import (
	"annie-baby/ent/video"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoCreate is the builder for creating a Video entity.
type VideoCreate struct {
	config
	mutation *VideoMutation
	hooks    []Hook
}

// SetLink sets the "link" field.
func (vc *VideoCreate) SetLink(s string) *VideoCreate {
	vc.mutation.SetLink(s)
	return vc
}

// SetType sets the "type" field.
func (vc *VideoCreate) SetType(s string) *VideoCreate {
	vc.mutation.SetType(s)
	return vc
}

// SetTitle sets the "title" field.
func (vc *VideoCreate) SetTitle(s string) *VideoCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetOutputDir sets the "output_dir" field.
func (vc *VideoCreate) SetOutputDir(s string) *VideoCreate {
	vc.mutation.SetOutputDir(s)
	return vc
}

// SetProxy sets the "proxy" field.
func (vc *VideoCreate) SetProxy(s string) *VideoCreate {
	vc.mutation.SetProxy(s)
	return vc
}

// SetStatus sets the "status" field.
func (vc *VideoCreate) SetStatus(v video.Status) *VideoCreate {
	vc.mutation.SetStatus(v)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VideoCreate) SetCreatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableCreatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VideoCreate) SetUpdatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableUpdatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// Mutation returns the VideoMutation object of the builder.
func (vc *VideoCreate) Mutation() *VideoMutation {
	return vc.mutation
}

// Save creates the Video in the database.
func (vc *VideoCreate) Save(ctx context.Context) (*Video, error) {
	var (
		err  error
		node *Video
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideoCreate) SaveX(ctx context.Context) *Video {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideoCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideoCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VideoCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := video.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := video.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideoCreate) check() error {
	if _, ok := vc.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`ent: missing required field "link"`)}
	}
	if v, ok := vc.mutation.Link(); ok {
		if err := video.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "link": %w`, err)}
		}
	}
	if _, ok := vc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := vc.mutation.GetType(); ok {
		if err := video.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := vc.mutation.Title(); ok {
		if err := video.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := vc.mutation.OutputDir(); !ok {
		return &ValidationError{Name: "output_dir", err: errors.New(`ent: missing required field "output_dir"`)}
	}
	if _, ok := vc.mutation.Proxy(); !ok {
		return &ValidationError{Name: "proxy", err: errors.New(`ent: missing required field "proxy"`)}
	}
	if _, ok := vc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := vc.mutation.Status(); ok {
		if err := video.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (vc *VideoCreate) sqlSave(ctx context.Context) (*Video, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VideoCreate) createSpec() (*Video, *sqlgraph.CreateSpec) {
	var (
		_node = &Video{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: video.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: video.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Link(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldLink,
		})
		_node.Link = value
	}
	if value, ok := vc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldType,
		})
		_node.Type = value
	}
	if value, ok := vc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := vc.mutation.OutputDir(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldOutputDir,
		})
		_node.OutputDir = value
	}
	if value, ok := vc.mutation.Proxy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldProxy,
		})
		_node.Proxy = value
	}
	if value, ok := vc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: video.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: video.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: video.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// VideoCreateBulk is the builder for creating many Video entities in bulk.
type VideoCreateBulk struct {
	config
	builders []*VideoCreate
}

// Save creates the Video entities in the database.
func (vcb *VideoCreateBulk) Save(ctx context.Context) ([]*Video, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Video, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideoCreateBulk) SaveX(ctx context.Context) []*Video {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideoCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideoCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
