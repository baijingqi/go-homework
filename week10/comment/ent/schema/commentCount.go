package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
    "time"
)

type CommentCount struct {
	ent.Schema
}

func (CommentCount) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "comments_count"},
    }
}

// Fields of the CommentCount.
func (CommentCount) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("id").Positive(),
        field.Uint32("praise_num").Default(0),
        field.Uint32("reply_num").Default(0),
        field.Uint32("dislike_num").Default(0),

        field.Time("created_at").
            Default(time.Now),
        field.Time("updated_at").
            Default(time.Now),
    }
}

// Edges of the Comment.
func (CommentCount) Edges() []ent.Edge {
	return nil
}
