// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/user"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/userprovider"
	"github.com/google/uuid"
)

// UserProviderCreate is the builder for creating a UserProvider entity.
type UserProviderCreate struct {
	config
	mutation *UserProviderMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (upc *UserProviderCreate) SetUserID(u uuid.UUID) *UserProviderCreate {
	upc.mutation.SetUserID(u)
	return upc
}

// SetPicture sets the "picture" field.
func (upc *UserProviderCreate) SetPicture(s string) *UserProviderCreate {
	upc.mutation.SetPicture(s)
	return upc
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (upc *UserProviderCreate) SetNillablePicture(s *string) *UserProviderCreate {
	if s != nil {
		upc.SetPicture(*s)
	}
	return upc
}

// SetName sets the "name" field.
func (upc *UserProviderCreate) SetName(u userprovider.Name) *UserProviderCreate {
	upc.mutation.SetName(u)
	return upc
}

// SetCreatedAt sets the "createdAt" field.
func (upc *UserProviderCreate) SetCreatedAt(t time.Time) *UserProviderCreate {
	upc.mutation.SetCreatedAt(t)
	return upc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (upc *UserProviderCreate) SetNillableCreatedAt(t *time.Time) *UserProviderCreate {
	if t != nil {
		upc.SetCreatedAt(*t)
	}
	return upc
}

// SetUpdatedAt sets the "updatedAt" field.
func (upc *UserProviderCreate) SetUpdatedAt(t time.Time) *UserProviderCreate {
	upc.mutation.SetUpdatedAt(t)
	return upc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (upc *UserProviderCreate) SetNillableUpdatedAt(t *time.Time) *UserProviderCreate {
	if t != nil {
		upc.SetUpdatedAt(*t)
	}
	return upc
}

// SetUser sets the "user" edge to the User entity.
func (upc *UserProviderCreate) SetUser(u *User) *UserProviderCreate {
	return upc.SetUserID(u.ID)
}

// Mutation returns the UserProviderMutation object of the builder.
func (upc *UserProviderCreate) Mutation() *UserProviderMutation {
	return upc.mutation
}

// Save creates the UserProvider in the database.
func (upc *UserProviderCreate) Save(ctx context.Context) (*UserProvider, error) {
	upc.defaults()
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserProviderCreate) SaveX(ctx context.Context) *UserProvider {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserProviderCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserProviderCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UserProviderCreate) defaults() {
	if _, ok := upc.mutation.CreatedAt(); !ok {
		v := userprovider.DefaultCreatedAt()
		upc.mutation.SetCreatedAt(v)
	}
	if _, ok := upc.mutation.UpdatedAt(); !ok {
		v := userprovider.DefaultUpdatedAt()
		upc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserProviderCreate) check() error {
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserProvider.user_id"`)}
	}
	if _, ok := upc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserProvider.name"`)}
	}
	if v, ok := upc.mutation.Name(); ok {
		if err := userprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserProvider.name": %w`, err)}
		}
	}
	if _, ok := upc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "UserProvider.createdAt"`)}
	}
	if _, ok := upc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "UserProvider.updatedAt"`)}
	}
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserProvider.user"`)}
	}
	return nil
}

func (upc *UserProviderCreate) sqlSave(ctx context.Context) (*UserProvider, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UserProviderCreate) createSpec() (*UserProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &UserProvider{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userprovider.Table, sqlgraph.NewFieldSpec(userprovider.FieldID, field.TypeInt))
	)
	if value, ok := upc.mutation.Picture(); ok {
		_spec.SetField(userprovider.FieldPicture, field.TypeString, value)
		_node.Picture = value
	}
	if value, ok := upc.mutation.Name(); ok {
		_spec.SetField(userprovider.FieldName, field.TypeEnum, value)
		_node.Name = value
	}
	if value, ok := upc.mutation.CreatedAt(); ok {
		_spec.SetField(userprovider.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := upc.mutation.UpdatedAt(); ok {
		_spec.SetField(userprovider.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := upc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userprovider.UserTable,
			Columns: []string{userprovider.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserProviderCreateBulk is the builder for creating many UserProvider entities in bulk.
type UserProviderCreateBulk struct {
	config
	err      error
	builders []*UserProviderCreate
}

// Save creates the UserProvider entities in the database.
func (upcb *UserProviderCreateBulk) Save(ctx context.Context) ([]*UserProvider, error) {
	if upcb.err != nil {
		return nil, upcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserProvider, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserProviderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserProviderCreateBulk) SaveX(ctx context.Context) []*UserProvider {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserProviderCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
