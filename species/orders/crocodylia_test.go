package orders

import "testing"

func TestCrocodylia(t *testing.T) {
	var s string = "Crocodylia"
	if ok := crocodylia == s; !ok {
		t.Fatalf("crocodylia != %s", s)
	}
}
