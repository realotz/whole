package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").NotEmpty().Unique().Comment("名称"),
		field.String("service").NotEmpty().Unique().Comment("服务"),
		field.String("path").NotEmpty().Unique().Comment("路径"),
		field.String("action").NotEmpty().Unique().Comment("动作"),
		field.String("desc").Optional().Comment("备注"),
	}
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Customer.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("permissions"),
	}
}
