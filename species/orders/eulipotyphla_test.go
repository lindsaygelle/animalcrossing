package orders

import "testing"

func TestEulipotyphla(t *testing.T) {
	var s string = "Eulipotyphla"
	if ok := eulipotyphla == s; !ok {
		t.Fatalf("eulipotyphla != %s", eulipotyphla)
	}
}
