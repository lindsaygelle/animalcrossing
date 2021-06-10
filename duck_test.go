package animalcrossing

import "testing"

func TestDuckName(t *testing.T) {
	if ok := Duck.Name() == duck; !ok {
		t.Fatalf("Duck.Name() != duck")
	}
}
