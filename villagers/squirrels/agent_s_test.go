package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAgentSAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := AgentS.Animal() == s; !ok {
		t.Fatalf("%s != %s", AgentS.Animal(), s)
	}
}

func TestAgentSName(t *testing.T) {
	if ok := AgentS.Name() == agentS; !ok {
		t.Fatalf("%s != %s", AgentS.Name(), agentS)
	}
}
