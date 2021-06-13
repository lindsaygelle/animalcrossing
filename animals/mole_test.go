package animals

import "testing"

func TestMoleName(t *testing.T) {
	if ok := Mole.Name() == mole; !ok {
		t.Fatal("Mole.Name() != mole")
	}
}
