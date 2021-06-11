package animalcrossing

import "testing"

func TestPeacockName(t *testing.T) {
	if ok := Peacock.Name() == peacock; !ok {
		t.Fatal("Peacock.Name() != peacock")
	}
}
