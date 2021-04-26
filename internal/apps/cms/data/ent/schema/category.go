package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Category struct{ ent.Schema }

// todo 请检查数据格式，并且去除必要参数的Optional方法
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment(""),
		field.String("name").Optional().Comment("分类名称"),
		field.Int64("pid").Optional().Comment("父级分类id"),
		field.String("icon").Optional().Comment("分类图标"),
		field.String("desc").Optional().Comment("分类简介"),
	}
}

func (Category) Mixin() []ent.Mixin { return []ent.Mixin{mixin.Time{}} }
