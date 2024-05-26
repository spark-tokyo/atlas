//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	seed := "./schema"
	option := &gen.Config{}
	if err := entc.Generate(seed, option); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
