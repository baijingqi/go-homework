// Code generated by entc, DO NOT EDIT.

package ent

import (
	"comment/ent/commentcount"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CommentCount is the model entity for the CommentCount schema.
type CommentCount struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// PraiseNum holds the value of the "praise_num" field.
	PraiseNum uint32 `json:"praise_num,omitempty"`
	// ReplyNum holds the value of the "reply_num" field.
	ReplyNum uint32 `json:"reply_num,omitempty"`
	// DislikeNum holds the value of the "dislike_num" field.
	DislikeNum uint32 `json:"dislike_num,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CommentCount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case commentcount.FieldID, commentcount.FieldPraiseNum, commentcount.FieldReplyNum, commentcount.FieldDislikeNum:
			values[i] = new(sql.NullInt64)
		case commentcount.FieldCreatedAt, commentcount.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CommentCount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CommentCount fields.
func (cc *CommentCount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case commentcount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = uint64(value.Int64)
		case commentcount.FieldPraiseNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field praise_num", values[i])
			} else if value.Valid {
				cc.PraiseNum = uint32(value.Int64)
			}
		case commentcount.FieldReplyNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reply_num", values[i])
			} else if value.Valid {
				cc.ReplyNum = uint32(value.Int64)
			}
		case commentcount.FieldDislikeNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dislike_num", values[i])
			} else if value.Valid {
				cc.DislikeNum = uint32(value.Int64)
			}
		case commentcount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cc.CreatedAt = value.Time
			}
		case commentcount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cc.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CommentCount.
// Note that you need to call CommentCount.Unwrap() before calling this method if this CommentCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *CommentCount) Update() *CommentCountUpdateOne {
	return (&CommentCountClient{config: cc.config}).UpdateOne(cc)
}

// Unwrap unwraps the CommentCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *CommentCount) Unwrap() *CommentCount {
	tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CommentCount is not a transactional entity")
	}
	cc.config.driver = tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *CommentCount) String() string {
	var builder strings.Builder
	builder.WriteString("CommentCount(")
	builder.WriteString(fmt.Sprintf("id=%v", cc.ID))
	builder.WriteString(", praise_num=")
	builder.WriteString(fmt.Sprintf("%v", cc.PraiseNum))
	builder.WriteString(", reply_num=")
	builder.WriteString(fmt.Sprintf("%v", cc.ReplyNum))
	builder.WriteString(", dislike_num=")
	builder.WriteString(fmt.Sprintf("%v", cc.DislikeNum))
	builder.WriteString(", created_at=")
	builder.WriteString(cc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CommentCounts is a parsable slice of CommentCount.
type CommentCounts []*CommentCount

func (cc CommentCounts) config(cfg config) {
	for _i := range cc {
		cc[_i].config = cfg
	}
}