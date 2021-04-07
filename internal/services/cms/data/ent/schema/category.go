package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("pid").Comment("父级分类id"),
		field.String("name").Comment("分类名称"),
	}
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
