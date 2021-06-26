package main

import (
	"fmt"

	"github.com/lindsaygelle/animalcrossing/villagers"
	"github.com/lindsaygelle/animalcrossing/villagers/ghosts"
)

func villager(v villagers.Villager) {
	fmt.Println(v.Name())
}

func main() {
	villager(ghosts.Wisp)
}
