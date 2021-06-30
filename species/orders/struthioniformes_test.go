package orders

import "testing"

func TestStruthioniformes(t *testing.T) {
	var s string = "Struthioniformes"
	if ok := struthioniformes == s; !ok {
		t.Fatalf("struthioniformes != %s", s)
	}
}
