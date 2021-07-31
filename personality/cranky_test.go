package personality

import "testing"

// TestCrankyId tests whether Cranky has the correct ID.
func TestCrankyId(t *testing.T) {
	if v := Cranky.Id(); !(v == crankyId) {
		t.Fatalf("%s != %s", v, crankyId)
	}
}
