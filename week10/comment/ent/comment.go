// Code generated by entc, DO NOT EDIT.

package ent

import (
	"comment/ent/comment"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID uint64 `json:"uid,omitempty"`
	// RelationID holds the value of the "relation_id" field.
	RelationID uint64 `json:"relation_id,omitempty"`
	// ReplyToUID holds the value of the "reply_to_uid" field.
	ReplyToUID uint64 `json:"reply_to_uid,omitempty"`
	// RelationType holds the value of the "relation_type" field.
	RelationType uint32 `json:"relation_type,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID uint64 `json:"parent_id,omitempty"`
	// BelongCommentID holds the value of the "belong_comment_id" field.
	BelongCommentID uint64 `json:"belong_comment_id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldUID, comment.FieldRelationID, comment.FieldReplyToUID, comment.FieldRelationType, comment.FieldParentID, comment.FieldBelongCommentID:
			values[i] = new(sql.NullInt64)
		case comment.FieldContent:
			values[i] = new(sql.NullString)
		case comment.FieldCreatedAt, comment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case comment.FieldUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				c.UID = uint64(value.Int64)
			}
		case comment.FieldRelationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relation_id", values[i])
			} else if value.Valid {
				c.RelationID = uint64(value.Int64)
			}
		case comment.FieldReplyToUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reply_to_uid", values[i])
			} else if value.Valid {
				c.ReplyToUID = uint64(value.Int64)
			}
		case comment.FieldRelationType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relation_type", values[i])
			} else if value.Valid {
				c.RelationType = uint32(value.Int64)
			}
		case comment.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				c.ParentID = uint64(value.Int64)
			}
		case comment.FieldBelongCommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field belong_comment_id", values[i])
			} else if value.Valid {
				c.BelongCommentID = uint64(value.Int64)
			}
		case comment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", uid=")
	builder.WriteString(fmt.Sprintf("%v", c.UID))
	builder.WriteString(", relation_id=")
	builder.WriteString(fmt.Sprintf("%v", c.RelationID))
	builder.WriteString(", reply_to_uid=")
	builder.WriteString(fmt.Sprintf("%v", c.ReplyToUID))
	builder.WriteString(", relation_type=")
	builder.WriteString(fmt.Sprintf("%v", c.RelationType))
	builder.WriteString(", parent_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ParentID))
	builder.WriteString(", belong_comment_id=")
	builder.WriteString(fmt.Sprintf("%v", c.BelongCommentID))
	builder.WriteString(", content=")
	builder.WriteString(c.Content)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
