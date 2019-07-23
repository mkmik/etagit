package main

import (
	"log"
	"math/rand"
	"time"

	regen "github.com/zach-klippenstein/goregen"
)

var (
	etagGen = mustGenerator(regen.NewGenerator(
		"[a-zA-Z0-9]{12}",
		&regen.GeneratorArgs{
			RngSource: rand.NewSource(time.Now().UnixNano()),
		},
	))
)

func mustGenerator(g regen.Generator, err error) regen.Generator {
	if err != nil {
		log.Fatal(err)
	}
	return g
}