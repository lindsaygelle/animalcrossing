package animals

import "testing"

func TestChameleonName(t *testing.T) {
	if ok := Chameleon.Name() == chameleon; !ok {
		t.Fatal("Chameleon.Name() != chameleon")
	}
}
