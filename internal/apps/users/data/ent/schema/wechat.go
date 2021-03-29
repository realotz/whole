package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Wechat holds the schema definition for the Wechat entity.
type Wechat struct {
	ent.Schema
}

// Fields of the Wechat.
func (Wechat) Fields() []ent.Field {
	return []ent.Field{
		field.String("openid").NotEmpty().Comment("openid"),
		field.String("unionId").NotEmpty().Comment("unionId"),
		field.Enum("app_type").Values("weapp", "official_account").Comment("类型 小程序 公众号"),
		field.Bytes("meta_data").Optional().Comment("微信原始数据"),
	}
}

func (Wechat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Wechat.
func (Wechat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Member.Type).Ref("wechats").Unique(),
	}
}
