package orders

import "testing"

func TestRodentia(t *testing.T) {
	var s string = "Rodentia"
	if ok := rodentia == s; !ok {
		t.Fatalf("rodentia != %s", s)
	}
}
