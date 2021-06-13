package animals

import "testing"

func TestGoatName(t *testing.T) {
	if ok := Goat.Name() == goat; !ok {
		t.Fatal("Goat.Name() != goat")
	}
}
