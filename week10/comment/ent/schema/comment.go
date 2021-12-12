package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
    "time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "comments"},
    }
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("id").
            Positive(),
        field.Uint64("uid").
            Positive(),
        field.Uint64("relation_id").
            Positive().
            Min(1),
        field.Uint64("reply_to_uid").
            Default(0),

        field.Uint32("relation_type").
            Positive(),
        field.Uint64("parent_id").
            Default(0).
            Min(0),
        field.Uint64("belong_comment_id").
            Default(0),
        field.String("content").
            Default(""),
        field.Time("created_at").
            Default(time.Now),
        field.Time("updated_at").
            Default(time.Now),
    }
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
