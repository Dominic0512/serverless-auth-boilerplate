// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/predicate"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/user"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/userprovider"
	"github.com/google/uuid"
)

// UserProviderUpdate is the builder for updating UserProvider entities.
type UserProviderUpdate struct {
	config
	hooks    []Hook
	mutation *UserProviderMutation
}

// Where appends a list predicates to the UserProviderUpdate builder.
func (upu *UserProviderUpdate) Where(ps ...predicate.UserProvider) *UserProviderUpdate {
	upu.mutation.Where(ps...)
	return upu
}

// SetUserID sets the "user_id" field.
func (upu *UserProviderUpdate) SetUserID(u uuid.UUID) *UserProviderUpdate {
	upu.mutation.SetUserID(u)
	return upu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upu *UserProviderUpdate) SetNillableUserID(u *uuid.UUID) *UserProviderUpdate {
	if u != nil {
		upu.SetUserID(*u)
	}
	return upu
}

// SetPicture sets the "picture" field.
func (upu *UserProviderUpdate) SetPicture(s string) *UserProviderUpdate {
	upu.mutation.SetPicture(s)
	return upu
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (upu *UserProviderUpdate) SetNillablePicture(s *string) *UserProviderUpdate {
	if s != nil {
		upu.SetPicture(*s)
	}
	return upu
}

// ClearPicture clears the value of the "picture" field.
func (upu *UserProviderUpdate) ClearPicture() *UserProviderUpdate {
	upu.mutation.ClearPicture()
	return upu
}

// SetName sets the "name" field.
func (upu *UserProviderUpdate) SetName(u userprovider.Name) *UserProviderUpdate {
	upu.mutation.SetName(u)
	return upu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (upu *UserProviderUpdate) SetNillableName(u *userprovider.Name) *UserProviderUpdate {
	if u != nil {
		upu.SetName(*u)
	}
	return upu
}

// SetCreatedAt sets the "createdAt" field.
func (upu *UserProviderUpdate) SetCreatedAt(t time.Time) *UserProviderUpdate {
	upu.mutation.SetCreatedAt(t)
	return upu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (upu *UserProviderUpdate) SetNillableCreatedAt(t *time.Time) *UserProviderUpdate {
	if t != nil {
		upu.SetCreatedAt(*t)
	}
	return upu
}

// SetUpdatedAt sets the "updatedAt" field.
func (upu *UserProviderUpdate) SetUpdatedAt(t time.Time) *UserProviderUpdate {
	upu.mutation.SetUpdatedAt(t)
	return upu
}

// SetUser sets the "user" edge to the User entity.
func (upu *UserProviderUpdate) SetUser(u *User) *UserProviderUpdate {
	return upu.SetUserID(u.ID)
}

// Mutation returns the UserProviderMutation object of the builder.
func (upu *UserProviderUpdate) Mutation() *UserProviderMutation {
	return upu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upu *UserProviderUpdate) ClearUser() *UserProviderUpdate {
	upu.mutation.ClearUser()
	return upu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UserProviderUpdate) Save(ctx context.Context) (int, error) {
	upu.defaults()
	return withHooks(ctx, upu.sqlSave, upu.mutation, upu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UserProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UserProviderUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UserProviderUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upu *UserProviderUpdate) defaults() {
	if _, ok := upu.mutation.UpdatedAt(); !ok {
		v := userprovider.UpdateDefaultUpdatedAt()
		upu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upu *UserProviderUpdate) check() error {
	if v, ok := upu.mutation.Name(); ok {
		if err := userprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserProvider.name": %w`, err)}
		}
	}
	if _, ok := upu.mutation.UserID(); upu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserProvider.user"`)
	}
	return nil
}

func (upu *UserProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := upu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userprovider.Table, userprovider.Columns, sqlgraph.NewFieldSpec(userprovider.FieldID, field.TypeInt))
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.Picture(); ok {
		_spec.SetField(userprovider.FieldPicture, field.TypeString, value)
	}
	if upu.mutation.PictureCleared() {
		_spec.ClearField(userprovider.FieldPicture, field.TypeString)
	}
	if value, ok := upu.mutation.Name(); ok {
		_spec.SetField(userprovider.FieldName, field.TypeEnum, value)
	}
	if value, ok := upu.mutation.CreatedAt(); ok {
		_spec.SetField(userprovider.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := upu.mutation.UpdatedAt(); ok {
		_spec.SetField(userprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if upu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	upu.mutation.done = true
	return n, nil
}

// UserProviderUpdateOne is the builder for updating a single UserProvider entity.
type UserProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserProviderMutation
}

// SetUserID sets the "user_id" field.
func (upuo *UserProviderUpdateOne) SetUserID(u uuid.UUID) *UserProviderUpdateOne {
	upuo.mutation.SetUserID(u)
	return upuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upuo *UserProviderUpdateOne) SetNillableUserID(u *uuid.UUID) *UserProviderUpdateOne {
	if u != nil {
		upuo.SetUserID(*u)
	}
	return upuo
}

// SetPicture sets the "picture" field.
func (upuo *UserProviderUpdateOne) SetPicture(s string) *UserProviderUpdateOne {
	upuo.mutation.SetPicture(s)
	return upuo
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (upuo *UserProviderUpdateOne) SetNillablePicture(s *string) *UserProviderUpdateOne {
	if s != nil {
		upuo.SetPicture(*s)
	}
	return upuo
}

// ClearPicture clears the value of the "picture" field.
func (upuo *UserProviderUpdateOne) ClearPicture() *UserProviderUpdateOne {
	upuo.mutation.ClearPicture()
	return upuo
}

// SetName sets the "name" field.
func (upuo *UserProviderUpdateOne) SetName(u userprovider.Name) *UserProviderUpdateOne {
	upuo.mutation.SetName(u)
	return upuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (upuo *UserProviderUpdateOne) SetNillableName(u *userprovider.Name) *UserProviderUpdateOne {
	if u != nil {
		upuo.SetName(*u)
	}
	return upuo
}

// SetCreatedAt sets the "createdAt" field.
func (upuo *UserProviderUpdateOne) SetCreatedAt(t time.Time) *UserProviderUpdateOne {
	upuo.mutation.SetCreatedAt(t)
	return upuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (upuo *UserProviderUpdateOne) SetNillableCreatedAt(t *time.Time) *UserProviderUpdateOne {
	if t != nil {
		upuo.SetCreatedAt(*t)
	}
	return upuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (upuo *UserProviderUpdateOne) SetUpdatedAt(t time.Time) *UserProviderUpdateOne {
	upuo.mutation.SetUpdatedAt(t)
	return upuo
}

// SetUser sets the "user" edge to the User entity.
func (upuo *UserProviderUpdateOne) SetUser(u *User) *UserProviderUpdateOne {
	return upuo.SetUserID(u.ID)
}

// Mutation returns the UserProviderMutation object of the builder.
func (upuo *UserProviderUpdateOne) Mutation() *UserProviderMutation {
	return upuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upuo *UserProviderUpdateOne) ClearUser() *UserProviderUpdateOne {
	upuo.mutation.ClearUser()
	return upuo
}

// Where appends a list predicates to the UserProviderUpdate builder.
func (upuo *UserProviderUpdateOne) Where(ps ...predicate.UserProvider) *UserProviderUpdateOne {
	upuo.mutation.Where(ps...)
	return upuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upuo *UserProviderUpdateOne) Select(field string, fields ...string) *UserProviderUpdateOne {
	upuo.fields = append([]string{field}, fields...)
	return upuo
}

// Save executes the query and returns the updated UserProvider entity.
func (upuo *UserProviderUpdateOne) Save(ctx context.Context) (*UserProvider, error) {
	upuo.defaults()
	return withHooks(ctx, upuo.sqlSave, upuo.mutation, upuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UserProviderUpdateOne) SaveX(ctx context.Context) *UserProvider {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UserProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UserProviderUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upuo *UserProviderUpdateOne) defaults() {
	if _, ok := upuo.mutation.UpdatedAt(); !ok {
		v := userprovider.UpdateDefaultUpdatedAt()
		upuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upuo *UserProviderUpdateOne) check() error {
	if v, ok := upuo.mutation.Name(); ok {
		if err := userprovider.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UserProvider.name": %w`, err)}
		}
	}
	if _, ok := upuo.mutation.UserID(); upuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserProvider.user"`)
	}
	return nil
}

func (upuo *UserProviderUpdateOne) sqlSave(ctx context.Context) (_node *UserProvider, err error) {
	if err := upuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userprovider.Table, userprovider.Columns, sqlgraph.NewFieldSpec(userprovider.FieldID, field.TypeInt))
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := upuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userprovider.FieldID)
		for _, f := range fields {
			if !userprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upuo.mutation.Picture(); ok {
		_spec.SetField(userprovider.FieldPicture, field.TypeString, value)
	}
	if upuo.mutation.PictureCleared() {
		_spec.ClearField(userprovider.FieldPicture, field.TypeString)
	}
	if value, ok := upuo.mutation.Name(); ok {
		_spec.SetField(userprovider.FieldName, field.TypeEnum, value)
	}
	if value, ok := upuo.mutation.CreatedAt(); ok {
		_spec.SetField(userprovider.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := upuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if upuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserProvider{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	upuo.mutation.done = true
	return _node, nil
}
