package orders

import "testing"

func TestPerissodactyla(t *testing.T) {
	var s string = "Perissodactyla"
	if ok := perissodactyla == s; !ok {
		t.Fatalf("perissodactyla != %s", s)
	}
}
