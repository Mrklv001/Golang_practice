package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateCompliment(name string) string {
	comliment := []string{
		"Ты великолепен,",
		"У тебя потрясающая улыбка,",
		"Ты вдохновляешь,",
	}
	rand.Seed(time.Now().UnixNano())
	randomCompliment := comliment[rand.Intn(len(comliment))]

	return fmt.Sprintf("%s %s!", randomCompliment, name)
}
