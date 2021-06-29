package genuses

import "testing"

func TestGenusGiraffa(t *testing.T) {
	var s string = "Giraffa"
	if ok := giraffa == s; !ok {
		t.Fatalf("giraffa != %s", s)
	}
}
