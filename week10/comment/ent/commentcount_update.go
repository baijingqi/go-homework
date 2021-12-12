// Code generated by entc, DO NOT EDIT.

package ent

import (
	"comment/ent/commentcount"
	"comment/ent/predicate"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentCountUpdate is the builder for updating CommentCount entities.
type CommentCountUpdate struct {
	config
	hooks    []Hook
	mutation *CommentCountMutation
}

// Where appends a list predicates to the CommentCountUpdate builder.
func (ccu *CommentCountUpdate) Where(ps ...predicate.CommentCount) *CommentCountUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetPraiseNum sets the "praise_num" field.
func (ccu *CommentCountUpdate) SetPraiseNum(u uint32) *CommentCountUpdate {
	ccu.mutation.ResetPraiseNum()
	ccu.mutation.SetPraiseNum(u)
	return ccu
}

// SetNillablePraiseNum sets the "praise_num" field if the given value is not nil.
func (ccu *CommentCountUpdate) SetNillablePraiseNum(u *uint32) *CommentCountUpdate {
	if u != nil {
		ccu.SetPraiseNum(*u)
	}
	return ccu
}

// AddPraiseNum adds u to the "praise_num" field.
func (ccu *CommentCountUpdate) AddPraiseNum(u uint32) *CommentCountUpdate {
	ccu.mutation.AddPraiseNum(u)
	return ccu
}

// SetReplyNum sets the "reply_num" field.
func (ccu *CommentCountUpdate) SetReplyNum(u uint32) *CommentCountUpdate {
	ccu.mutation.ResetReplyNum()
	ccu.mutation.SetReplyNum(u)
	return ccu
}

// SetNillableReplyNum sets the "reply_num" field if the given value is not nil.
func (ccu *CommentCountUpdate) SetNillableReplyNum(u *uint32) *CommentCountUpdate {
	if u != nil {
		ccu.SetReplyNum(*u)
	}
	return ccu
}

// AddReplyNum adds u to the "reply_num" field.
func (ccu *CommentCountUpdate) AddReplyNum(u uint32) *CommentCountUpdate {
	ccu.mutation.AddReplyNum(u)
	return ccu
}

// SetDislikeNum sets the "dislike_num" field.
func (ccu *CommentCountUpdate) SetDislikeNum(u uint32) *CommentCountUpdate {
	ccu.mutation.ResetDislikeNum()
	ccu.mutation.SetDislikeNum(u)
	return ccu
}

// SetNillableDislikeNum sets the "dislike_num" field if the given value is not nil.
func (ccu *CommentCountUpdate) SetNillableDislikeNum(u *uint32) *CommentCountUpdate {
	if u != nil {
		ccu.SetDislikeNum(*u)
	}
	return ccu
}

// AddDislikeNum adds u to the "dislike_num" field.
func (ccu *CommentCountUpdate) AddDislikeNum(u uint32) *CommentCountUpdate {
	ccu.mutation.AddDislikeNum(u)
	return ccu
}

// SetCreatedAt sets the "created_at" field.
func (ccu *CommentCountUpdate) SetCreatedAt(t time.Time) *CommentCountUpdate {
	ccu.mutation.SetCreatedAt(t)
	return ccu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccu *CommentCountUpdate) SetNillableCreatedAt(t *time.Time) *CommentCountUpdate {
	if t != nil {
		ccu.SetCreatedAt(*t)
	}
	return ccu
}

// SetUpdatedAt sets the "updated_at" field.
func (ccu *CommentCountUpdate) SetUpdatedAt(t time.Time) *CommentCountUpdate {
	ccu.mutation.SetUpdatedAt(t)
	return ccu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccu *CommentCountUpdate) SetNillableUpdatedAt(t *time.Time) *CommentCountUpdate {
	if t != nil {
		ccu.SetUpdatedAt(*t)
	}
	return ccu
}

// Mutation returns the CommentCountMutation object of the builder.
func (ccu *CommentCountUpdate) Mutation() *CommentCountMutation {
	return ccu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CommentCountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ccu.hooks) == 0 {
		affected, err = ccu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccu.mutation = mutation
			affected, err = ccu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccu.hooks) - 1; i >= 0; i-- {
			if ccu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CommentCountUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CommentCountUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CommentCountUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ccu *CommentCountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commentcount.Table,
			Columns: commentcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: commentcount.FieldID,
			},
		},
	}
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.PraiseNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldPraiseNum,
		})
	}
	if value, ok := ccu.mutation.AddedPraiseNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldPraiseNum,
		})
	}
	if value, ok := ccu.mutation.ReplyNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldReplyNum,
		})
	}
	if value, ok := ccu.mutation.AddedReplyNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldReplyNum,
		})
	}
	if value, ok := ccu.mutation.DislikeNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldDislikeNum,
		})
	}
	if value, ok := ccu.mutation.AddedDislikeNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldDislikeNum,
		})
	}
	if value, ok := ccu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcount.FieldCreatedAt,
		})
	}
	if value, ok := ccu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcount.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commentcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CommentCountUpdateOne is the builder for updating a single CommentCount entity.
type CommentCountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentCountMutation
}

// SetPraiseNum sets the "praise_num" field.
func (ccuo *CommentCountUpdateOne) SetPraiseNum(u uint32) *CommentCountUpdateOne {
	ccuo.mutation.ResetPraiseNum()
	ccuo.mutation.SetPraiseNum(u)
	return ccuo
}

// SetNillablePraiseNum sets the "praise_num" field if the given value is not nil.
func (ccuo *CommentCountUpdateOne) SetNillablePraiseNum(u *uint32) *CommentCountUpdateOne {
	if u != nil {
		ccuo.SetPraiseNum(*u)
	}
	return ccuo
}

// AddPraiseNum adds u to the "praise_num" field.
func (ccuo *CommentCountUpdateOne) AddPraiseNum(u uint32) *CommentCountUpdateOne {
	ccuo.mutation.AddPraiseNum(u)
	return ccuo
}

// SetReplyNum sets the "reply_num" field.
func (ccuo *CommentCountUpdateOne) SetReplyNum(u uint32) *CommentCountUpdateOne {
	ccuo.mutation.ResetReplyNum()
	ccuo.mutation.SetReplyNum(u)
	return ccuo
}

// SetNillableReplyNum sets the "reply_num" field if the given value is not nil.
func (ccuo *CommentCountUpdateOne) SetNillableReplyNum(u *uint32) *CommentCountUpdateOne {
	if u != nil {
		ccuo.SetReplyNum(*u)
	}
	return ccuo
}

// AddReplyNum adds u to the "reply_num" field.
func (ccuo *CommentCountUpdateOne) AddReplyNum(u uint32) *CommentCountUpdateOne {
	ccuo.mutation.AddReplyNum(u)
	return ccuo
}

// SetDislikeNum sets the "dislike_num" field.
func (ccuo *CommentCountUpdateOne) SetDislikeNum(u uint32) *CommentCountUpdateOne {
	ccuo.mutation.ResetDislikeNum()
	ccuo.mutation.SetDislikeNum(u)
	return ccuo
}

// SetNillableDislikeNum sets the "dislike_num" field if the given value is not nil.
func (ccuo *CommentCountUpdateOne) SetNillableDislikeNum(u *uint32) *CommentCountUpdateOne {
	if u != nil {
		ccuo.SetDislikeNum(*u)
	}
	return ccuo
}

// AddDislikeNum adds u to the "dislike_num" field.
func (ccuo *CommentCountUpdateOne) AddDislikeNum(u uint32) *CommentCountUpdateOne {
	ccuo.mutation.AddDislikeNum(u)
	return ccuo
}

// SetCreatedAt sets the "created_at" field.
func (ccuo *CommentCountUpdateOne) SetCreatedAt(t time.Time) *CommentCountUpdateOne {
	ccuo.mutation.SetCreatedAt(t)
	return ccuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccuo *CommentCountUpdateOne) SetNillableCreatedAt(t *time.Time) *CommentCountUpdateOne {
	if t != nil {
		ccuo.SetCreatedAt(*t)
	}
	return ccuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ccuo *CommentCountUpdateOne) SetUpdatedAt(t time.Time) *CommentCountUpdateOne {
	ccuo.mutation.SetUpdatedAt(t)
	return ccuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccuo *CommentCountUpdateOne) SetNillableUpdatedAt(t *time.Time) *CommentCountUpdateOne {
	if t != nil {
		ccuo.SetUpdatedAt(*t)
	}
	return ccuo
}

// Mutation returns the CommentCountMutation object of the builder.
func (ccuo *CommentCountUpdateOne) Mutation() *CommentCountMutation {
	return ccuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CommentCountUpdateOne) Select(field string, fields ...string) *CommentCountUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CommentCount entity.
func (ccuo *CommentCountUpdateOne) Save(ctx context.Context) (*CommentCount, error) {
	var (
		err  error
		node *CommentCount
	)
	if len(ccuo.hooks) == 0 {
		node, err = ccuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentCountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccuo.mutation = mutation
			node, err = ccuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ccuo.hooks) - 1; i >= 0; i-- {
			if ccuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CommentCountUpdateOne) SaveX(ctx context.Context) *CommentCount {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CommentCountUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CommentCountUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ccuo *CommentCountUpdateOne) sqlSave(ctx context.Context) (_node *CommentCount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commentcount.Table,
			Columns: commentcount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: commentcount.FieldID,
			},
		},
	}
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CommentCount.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commentcount.FieldID)
		for _, f := range fields {
			if !commentcount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commentcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.PraiseNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldPraiseNum,
		})
	}
	if value, ok := ccuo.mutation.AddedPraiseNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldPraiseNum,
		})
	}
	if value, ok := ccuo.mutation.ReplyNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldReplyNum,
		})
	}
	if value, ok := ccuo.mutation.AddedReplyNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldReplyNum,
		})
	}
	if value, ok := ccuo.mutation.DislikeNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldDislikeNum,
		})
	}
	if value, ok := ccuo.mutation.AddedDislikeNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commentcount.FieldDislikeNum,
		})
	}
	if value, ok := ccuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcount.FieldCreatedAt,
		})
	}
	if value, ok := ccuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: commentcount.FieldUpdatedAt,
		})
	}
	_node = &CommentCount{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commentcount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}