package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// id は宣言しなくてもデフォルトで設定されるが、UUIDにしたいため宣言してオーバーライド
		field.String("id").Comment("システムで使うユーザーID").Unique(),
		field.Int("age").Comment("年齢"),
		field.String("name").Comment("本名"),
		field.String("nickname").Comment("ユーザーネーム"),
		field.String("Email").Comment("メールアドレス").
			Unique(),
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
