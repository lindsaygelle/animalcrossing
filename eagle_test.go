package animalcrossing

import "testing"

func TestEagleName(t *testing.T) {
	if ok := Eagle.Name() == eagle; !ok {
		t.Fatal("Eagle.Name() != eagle")
	}
}
