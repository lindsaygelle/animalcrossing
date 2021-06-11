package animalcrossing

import "testing"

func TestKangarooName(t *testing.T) {
	if ok := Kangaroo.Name() == kangaroo; !ok {
		t.Fatalf("Kangaroo.Name() != kangaroo")
	}
}
