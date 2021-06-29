package families

import "testing"

func TestRhinocerotidae(t *testing.T) {
	var s string = "Rhinocerotidae"
	if ok := rhinocerotidae == s; !ok {
		t.Fatalf("rhinocerotidae != %s", s)
	}
}
