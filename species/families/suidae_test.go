package families

import "testing"

func TestSuidae(t *testing.T) {
	var s string = "Suidae"
	if ok := suidae == s; !ok {
		t.Fatalf("suidae != %s", s)
	}
}
