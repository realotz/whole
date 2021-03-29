package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive(),
		field.String("name").NotEmpty().Comment("文件名称"),
		field.String("type").NotEmpty().Comment("文件类型"),
		field.String("md5").NotEmpty().Comment("md5值"),
		field.String("path").NotEmpty().Comment("路径或url"),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
