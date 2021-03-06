// Code generated by entc, DO NOT EDIT.

package comment

import (
	"time"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldRelationID holds the string denoting the relation_id field in the database.
	FieldRelationID = "relation_id"
	// FieldReplyToUID holds the string denoting the reply_to_uid field in the database.
	FieldReplyToUID = "reply_to_uid"
	// FieldRelationType holds the string denoting the relation_type field in the database.
	FieldRelationType = "relation_type"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldBelongCommentID holds the string denoting the belong_comment_id field in the database.
	FieldBelongCommentID = "belong_comment_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the comment in the database.
	Table = "comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldRelationID,
	FieldReplyToUID,
	FieldRelationType,
	FieldParentID,
	FieldBelongCommentID,
	FieldContent,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	UIDValidator func(uint64) error
	// RelationIDValidator is a validator for the "relation_id" field. It is called by the builders before save.
	RelationIDValidator func(uint64) error
	// DefaultReplyToUID holds the default value on creation for the "reply_to_uid" field.
	DefaultReplyToUID uint64
	// RelationTypeValidator is a validator for the "relation_type" field. It is called by the builders before save.
	RelationTypeValidator func(uint32) error
	// DefaultParentID holds the default value on creation for the "parent_id" field.
	DefaultParentID uint64
	// ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	ParentIDValidator func(uint64) error
	// DefaultBelongCommentID holds the default value on creation for the "belong_comment_id" field.
	DefaultBelongCommentID uint64
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)

func test()  {

}