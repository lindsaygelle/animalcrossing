package families

import "testing"

func TestFamilyElephantidae(t *testing.T) {
	var s string = "Elephantidae"
	if ok := elephantidae == s; !ok {
		t.Fatalf("elephantidae != %s", s)
	}
}
