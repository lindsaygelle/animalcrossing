package families

import "testing"

func TestFamilyTapiridae(t *testing.T) {
	var s string = "Tapiridae"
	if ok := tapiridae == s; !ok {
		t.Fatalf("tapiridae != %s", s)
	}
}
