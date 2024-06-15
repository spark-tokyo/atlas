package fixtures

import "github.com/spark-tokyo/atlas/ent"

func SetUsers(tx *ent.Tx) []*ent.UserCreate {
	var users []*ent.UserCreate
	users = append(users, tx.User.Create().
		SetID("userId_001").
		SetAge(2).
		SetName("name").
		SetNickname("nickname").
		SetEmail("email1"),
	)
	users = append(users, tx.User.Create().
		SetID("userId_002").
		SetAge(3).
		SetName("name2").
		SetNickname("nickname2").
		SetEmail("email2"),
	)
	return users
}
