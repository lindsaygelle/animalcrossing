package families

import "testing"

func TestMyrmecophagidae(t *testing.T) {
	var s string = "Myrmecophagidae"
	if ok := myrmecophagidae == s; !ok {
		t.Fatalf("myrmecophagidae != %s", s)
	}
}
