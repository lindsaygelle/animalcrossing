package animals

import "testing"

func TestHamsterName(t *testing.T) {
	if ok := Hamster.Name() == hamster; !ok {
		t.Fatal("Hamster.Name() != hamster")
	}
}
