// Code generated by entc, DO NOT EDIT.

package ent

import (
	"comment/ent/comment"
	"comment/ent/predicate"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUID sets the "uid" field.
func (cu *CommentUpdate) SetUID(u uint64) *CommentUpdate {
	cu.mutation.ResetUID()
	cu.mutation.SetUID(u)
	return cu
}

// AddUID adds u to the "uid" field.
func (cu *CommentUpdate) AddUID(u uint64) *CommentUpdate {
	cu.mutation.AddUID(u)
	return cu
}

// SetRelationID sets the "relation_id" field.
func (cu *CommentUpdate) SetRelationID(u uint64) *CommentUpdate {
	cu.mutation.ResetRelationID()
	cu.mutation.SetRelationID(u)
	return cu
}

// AddRelationID adds u to the "relation_id" field.
func (cu *CommentUpdate) AddRelationID(u uint64) *CommentUpdate {
	cu.mutation.AddRelationID(u)
	return cu
}

// SetReplyToUID sets the "reply_to_uid" field.
func (cu *CommentUpdate) SetReplyToUID(u uint64) *CommentUpdate {
	cu.mutation.ResetReplyToUID()
	cu.mutation.SetReplyToUID(u)
	return cu
}

// SetNillableReplyToUID sets the "reply_to_uid" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableReplyToUID(u *uint64) *CommentUpdate {
	if u != nil {
		cu.SetReplyToUID(*u)
	}
	return cu
}

// AddReplyToUID adds u to the "reply_to_uid" field.
func (cu *CommentUpdate) AddReplyToUID(u uint64) *CommentUpdate {
	cu.mutation.AddReplyToUID(u)
	return cu
}

// SetRelationType sets the "relation_type" field.
func (cu *CommentUpdate) SetRelationType(u uint32) *CommentUpdate {
	cu.mutation.ResetRelationType()
	cu.mutation.SetRelationType(u)
	return cu
}

// AddRelationType adds u to the "relation_type" field.
func (cu *CommentUpdate) AddRelationType(u uint32) *CommentUpdate {
	cu.mutation.AddRelationType(u)
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CommentUpdate) SetParentID(u uint64) *CommentUpdate {
	cu.mutation.ResetParentID()
	cu.mutation.SetParentID(u)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableParentID(u *uint64) *CommentUpdate {
	if u != nil {
		cu.SetParentID(*u)
	}
	return cu
}

// AddParentID adds u to the "parent_id" field.
func (cu *CommentUpdate) AddParentID(u uint64) *CommentUpdate {
	cu.mutation.AddParentID(u)
	return cu
}

// SetBelongCommentID sets the "belong_comment_id" field.
func (cu *CommentUpdate) SetBelongCommentID(u uint64) *CommentUpdate {
	cu.mutation.ResetBelongCommentID()
	cu.mutation.SetBelongCommentID(u)
	return cu
}

// SetNillableBelongCommentID sets the "belong_comment_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableBelongCommentID(u *uint64) *CommentUpdate {
	if u != nil {
		cu.SetBelongCommentID(*u)
	}
	return cu
}

// AddBelongCommentID adds u to the "belong_comment_id" field.
func (cu *CommentUpdate) AddBelongCommentID(u uint64) *CommentUpdate {
	cu.mutation.AddBelongCommentID(u)
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableContent(s *string) *CommentUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommentUpdate) SetCreatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableCreatedAt(t *time.Time) *CommentUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUpdatedAt(t *time.Time) *CommentUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommentUpdate) check() error {
	if v, ok := cu.mutation.UID(); ok {
		if err := comment.UIDValidator(v); err != nil {
			return &ValidationError{Name: "uid", err: fmt.Errorf("ent: validator failed for field \"uid\": %w", err)}
		}
	}
	if v, ok := cu.mutation.RelationID(); ok {
		if err := comment.RelationIDValidator(v); err != nil {
			return &ValidationError{Name: "relation_id", err: fmt.Errorf("ent: validator failed for field \"relation_id\": %w", err)}
		}
	}
	if v, ok := cu.mutation.RelationType(); ok {
		if err := comment.RelationTypeValidator(v); err != nil {
			return &ValidationError{Name: "relation_type", err: fmt.Errorf("ent: validator failed for field \"relation_type\": %w", err)}
		}
	}
	if v, ok := cu.mutation.ParentID(); ok {
		if err := comment.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf("ent: validator failed for field \"parent_id\": %w", err)}
		}
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldUID,
		})
	}
	if value, ok := cu.mutation.AddedUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldUID,
		})
	}
	if value, ok := cu.mutation.RelationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldRelationID,
		})
	}
	if value, ok := cu.mutation.AddedRelationID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldRelationID,
		})
	}
	if value, ok := cu.mutation.ReplyToUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldReplyToUID,
		})
	}
	if value, ok := cu.mutation.AddedReplyToUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldReplyToUID,
		})
	}
	if value, ok := cu.mutation.RelationType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldRelationType,
		})
	}
	if value, ok := cu.mutation.AddedRelationType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldRelationType,
		})
	}
	if value, ok := cu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldParentID,
		})
	}
	if value, ok := cu.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldParentID,
		})
	}
	if value, ok := cu.mutation.BelongCommentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldBelongCommentID,
		})
	}
	if value, ok := cu.mutation.AddedBelongCommentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldBelongCommentID,
		})
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetUID sets the "uid" field.
func (cuo *CommentUpdateOne) SetUID(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetUID()
	cuo.mutation.SetUID(u)
	return cuo
}

// AddUID adds u to the "uid" field.
func (cuo *CommentUpdateOne) AddUID(u uint64) *CommentUpdateOne {
	cuo.mutation.AddUID(u)
	return cuo
}

// SetRelationID sets the "relation_id" field.
func (cuo *CommentUpdateOne) SetRelationID(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetRelationID()
	cuo.mutation.SetRelationID(u)
	return cuo
}

// AddRelationID adds u to the "relation_id" field.
func (cuo *CommentUpdateOne) AddRelationID(u uint64) *CommentUpdateOne {
	cuo.mutation.AddRelationID(u)
	return cuo
}

// SetReplyToUID sets the "reply_to_uid" field.
func (cuo *CommentUpdateOne) SetReplyToUID(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetReplyToUID()
	cuo.mutation.SetReplyToUID(u)
	return cuo
}

// SetNillableReplyToUID sets the "reply_to_uid" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableReplyToUID(u *uint64) *CommentUpdateOne {
	if u != nil {
		cuo.SetReplyToUID(*u)
	}
	return cuo
}

// AddReplyToUID adds u to the "reply_to_uid" field.
func (cuo *CommentUpdateOne) AddReplyToUID(u uint64) *CommentUpdateOne {
	cuo.mutation.AddReplyToUID(u)
	return cuo
}

// SetRelationType sets the "relation_type" field.
func (cuo *CommentUpdateOne) SetRelationType(u uint32) *CommentUpdateOne {
	cuo.mutation.ResetRelationType()
	cuo.mutation.SetRelationType(u)
	return cuo
}

// AddRelationType adds u to the "relation_type" field.
func (cuo *CommentUpdateOne) AddRelationType(u uint32) *CommentUpdateOne {
	cuo.mutation.AddRelationType(u)
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CommentUpdateOne) SetParentID(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetParentID()
	cuo.mutation.SetParentID(u)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableParentID(u *uint64) *CommentUpdateOne {
	if u != nil {
		cuo.SetParentID(*u)
	}
	return cuo
}

// AddParentID adds u to the "parent_id" field.
func (cuo *CommentUpdateOne) AddParentID(u uint64) *CommentUpdateOne {
	cuo.mutation.AddParentID(u)
	return cuo
}

// SetBelongCommentID sets the "belong_comment_id" field.
func (cuo *CommentUpdateOne) SetBelongCommentID(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetBelongCommentID()
	cuo.mutation.SetBelongCommentID(u)
	return cuo
}

// SetNillableBelongCommentID sets the "belong_comment_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableBelongCommentID(u *uint64) *CommentUpdateOne {
	if u != nil {
		cuo.SetBelongCommentID(*u)
	}
	return cuo
}

// AddBelongCommentID adds u to the "belong_comment_id" field.
func (cuo *CommentUpdateOne) AddBelongCommentID(u uint64) *CommentUpdateOne {
	cuo.mutation.AddBelongCommentID(u)
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableContent(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommentUpdateOne) SetCreatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCreatedAt(t *time.Time) *CommentUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUpdatedAt(t *time.Time) *CommentUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommentUpdateOne) check() error {
	if v, ok := cuo.mutation.UID(); ok {
		if err := comment.UIDValidator(v); err != nil {
			return &ValidationError{Name: "uid", err: fmt.Errorf("ent: validator failed for field \"uid\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.RelationID(); ok {
		if err := comment.RelationIDValidator(v); err != nil {
			return &ValidationError{Name: "relation_id", err: fmt.Errorf("ent: validator failed for field \"relation_id\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.RelationType(); ok {
		if err := comment.RelationTypeValidator(v); err != nil {
			return &ValidationError{Name: "relation_type", err: fmt.Errorf("ent: validator failed for field \"relation_type\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.ParentID(); ok {
		if err := comment.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf("ent: validator failed for field \"parent_id\": %w", err)}
		}
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: comment.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Comment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldUID,
		})
	}
	if value, ok := cuo.mutation.AddedUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldUID,
		})
	}
	if value, ok := cuo.mutation.RelationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldRelationID,
		})
	}
	if value, ok := cuo.mutation.AddedRelationID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldRelationID,
		})
	}
	if value, ok := cuo.mutation.ReplyToUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldReplyToUID,
		})
	}
	if value, ok := cuo.mutation.AddedReplyToUID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldReplyToUID,
		})
	}
	if value, ok := cuo.mutation.RelationType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldRelationType,
		})
	}
	if value, ok := cuo.mutation.AddedRelationType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldRelationType,
		})
	}
	if value, ok := cuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldParentID,
		})
	}
	if value, ok := cuo.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldParentID,
		})
	}
	if value, ok := cuo.mutation.BelongCommentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldBelongCommentID,
		})
	}
	if value, ok := cuo.mutation.AddedBelongCommentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldBelongCommentID,
		})
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
