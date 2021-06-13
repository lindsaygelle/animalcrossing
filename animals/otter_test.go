package animals

import "testing"

func TestOtterName(t *testing.T) {
	if ok := Otter.Name() == otter; !ok {
		t.Fatal("Otter.Name() != otter")
	}
}
