package families

import "testing"

func TestAccipitridae(t *testing.T) {
	var s string = "Accipitridae"
	if ok := accipitridae == s; !ok {
		t.Fatalf("accipitridae != %s", s)
	}
}
