package fixtures

import "github.com/spark-tokyo/atlas/ent"

func SetPets(tx *ent.Tx) []*ent.PetCreate {
	var pets []*ent.PetCreate
	pets = append(pets, tx.Pet.Create())
	return pets
}
