package animals

import "testing"

func TestGorillaName(t *testing.T) {
	if ok := Gorilla.Name() == gorilla; !ok {
		t.Fatal("Gorilla.Name() != gorilla")
	}
}
