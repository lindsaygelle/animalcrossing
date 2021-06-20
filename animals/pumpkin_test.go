package animals

import "testing"

func TestPumpkinName(t *testing.T) {
	if ok := Pumpkin.Name() == pumpkin; !ok {
		t.Fatal("Pumpkin.Name() != pumpkin")
	}
}
