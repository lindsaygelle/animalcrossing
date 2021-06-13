package animals

import "testing"

func TestFrogName(t *testing.T) {
	if ok := Frog.Name() == frog; !ok {
		t.Fatal("Frog.Name() != frog")
	}
}
