package animals

import "testing"

func TestOstrichName(t *testing.T) {
	if ok := Ostrich.Name() == ostrich; !ok {
		t.Fatal("Ostrich.Name() != goat")
	}
}
