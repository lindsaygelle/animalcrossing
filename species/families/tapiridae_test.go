package families

import "testing"

func TestTapiridae(t *testing.T) {
	var s string = "Tapiridae"
	if ok := tapiridae == s; !ok {
		t.Fatalf("tapiridae != %s", s)
	}
}
