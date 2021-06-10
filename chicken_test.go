package animalcrossing

import "testing"

func TestChickenName(t *testing.T) {
	if ok := Chicken.Name() == chicken; !ok {
		t.Fatal("Chicken.Name() != chicken")
	}
}
