package fixtures

import "github.com/spark-tokyo/atlas/ent"

func SetUsers(tx *ent.Tx) []*ent.UserCreate {
	var users []*ent.UserCreate
	users = append(users, tx.User.Create().SetAge(2).SetName("name").SetNickname("nickname"))
	users = append(users, tx.User.Create().SetAge(3).SetName("name2").SetNickname("nickname2"))
	return users
}
