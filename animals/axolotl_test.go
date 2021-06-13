package animals

import "testing"

func TestAxolotlName(t *testing.T) {
	if ok := Axolotl.Name() == axolotl; !ok {
		t.Fatal("Axolotl.Name() != axolotl")
	}
}
