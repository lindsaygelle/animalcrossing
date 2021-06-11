package animalcrossing

import "testing"

func TestMonkeyName(t *testing.T) {
	if ok := Monkey.Name() == monkey; !ok {
		t.Fatal("Monkey.Name() != monkey")
	}
}
