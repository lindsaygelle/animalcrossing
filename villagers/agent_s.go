package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	agentS string = "Agent S"
)

var (
	// AgentS is an Animal Crossing villager.
	//
	// AgentS is a Squirrel.
	AgentS Villager = villager{
		animal: animals.Squirrel.Name(),
		name:   agentS}
)
