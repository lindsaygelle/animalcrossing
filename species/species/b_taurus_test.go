package species

import "testing"

func TestBTaurus(t *testing.T) {
	var s string = "B. taurus"
	if ok := bTaurus == s; !ok {
		t.Fatalf("bTaurus != %s", s)
	}
}
