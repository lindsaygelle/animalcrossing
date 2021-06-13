package animals

import "testing"

func TestSheepName(t *testing.T) {
	if ok := Sheep.Name() == sheep; !ok {
		t.Fatal("Sheep.Name() != sheep")
	}
}
