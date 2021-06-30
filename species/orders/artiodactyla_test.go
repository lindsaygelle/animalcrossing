package orders

import "testing"

func TestArtiodactyla(t *testing.T) {
	var s string = "Artiodactyla"
	if ok := artiodactyla == s; !ok {
		t.Fatalf("artiodactyla != %s", s)
	}
}
