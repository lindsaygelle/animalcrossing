package families

import "testing"

func TestFamilyStruthionidae(t *testing.T) {
	var s string = "Struthionidae"
	if ok := struthionidae == s; !ok {
		t.Fatalf("struthionidae != %s", s)
	}
}
