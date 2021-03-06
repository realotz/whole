package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Enum("sex").Values("unknown", "man", "woman").Default("unknown").Comment("性别"),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Comment("uuid"),
		field.String("account").NotEmpty().Unique().Comment(" 账号"),
		field.String("avatar").Optional().Comment(" avatar"),
		field.String("name").Optional().Default("unknown").Comment("名称"),
		field.String("nick_name").Optional().Default("unknown").Comment("昵称"),
		field.String("email").Comment("邮箱"),
		field.String("mobile").Optional().Comment("手机号"),
		field.String("id_card").Optional().Comment("身份证"),
		field.Time("birthday").Optional().Comment("生日"),
		field.String("password").NotEmpty().Comment("密码"),
		field.String("salt").NotEmpty().Comment("加密盐"),
		field.String("last_ip").Optional().Comment("最后登陆ip"),
		field.Time("last_time").Optional().Comment("最后登陆时间"),
	}
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("wechats", Wechat.Type),
	}
}
