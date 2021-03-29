// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MembersColumns holds the columns for the "members" table.
	MembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "sex", Type: field.TypeEnum, Enums: []string{"unknown", "man", "woman"}, Default: "unknown"},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "account", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: "unknown"},
		{Name: "role", Type: field.TypeString, Nullable: true, Default: "member"},
		{Name: "nick_name", Type: field.TypeString, Nullable: true, Default: "unknown"},
		{Name: "email", Type: field.TypeString},
		{Name: "mobile", Type: field.TypeString, Nullable: true},
		{Name: "id_card", Type: field.TypeString, Nullable: true},
		{Name: "birthday", Type: field.TypeTime, Nullable: true},
		{Name: "password", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "last_ip", Type: field.TypeString, Nullable: true},
		{Name: "last_time", Type: field.TypeTime, Nullable: true},
	}
	// MembersTable holds the schema information for the "members" table.
	MembersTable = &schema.Table{
		Name:        "members",
		Columns:     MembersColumns,
		PrimaryKey:  []*schema.Column{MembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// WechatsColumns holds the columns for the "wechats" table.
	WechatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "openid", Type: field.TypeString},
		{Name: "union_id", Type: field.TypeString},
		{Name: "app_type", Type: field.TypeEnum, Enums: []string{"weapp", "official_account"}},
		{Name: "meta_data", Type: field.TypeBytes, Nullable: true},
		{Name: "member_wechats", Type: field.TypeInt64, Nullable: true},
	}
	// WechatsTable holds the schema information for the "wechats" table.
	WechatsTable = &schema.Table{
		Name:       "wechats",
		Columns:    WechatsColumns,
		PrimaryKey: []*schema.Column{WechatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "wechats_members_wechats",
				Columns:    []*schema.Column{WechatsColumns[7]},
				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MembersTable,
		WechatsTable,
	}
)

func init() {
	WechatsTable.ForeignKeys[0].RefTable = MembersTable
}
