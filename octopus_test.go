package animalcrossing

import "testing"

func TestOctopusName(t *testing.T) {
	if ok := Octopus.Name() == octopus; !ok {
		t.Fatal("Octopus.Name() != octopus")
	}
}
