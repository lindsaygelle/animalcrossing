package personality

import "testing"

// TestSisterlyId tests whether Sisterly has the correct ID.
func TestSisterlyId(t *testing.T) {
	if v := Sisterly.Id(); !(v == sisterlyId) {
		t.Fatalf("%s != %s", v, sisterlyId)
	}
}
