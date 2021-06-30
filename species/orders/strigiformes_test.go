package orders

import "testing"

func TestStrigiformes(t *testing.T) {
	var s string = "Strigiformes"
	if ok := strigiformes == s; !ok {
		t.Fatalf("strigiformes != %s", s)
	}
}
