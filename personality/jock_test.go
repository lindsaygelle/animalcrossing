package personality

import "testing"

// TestJockId tests whether Jock has the correct ID.
func TestJockId(t *testing.T) {
	if v := Jock.Id(); !(v == jockId) {
		t.Fatalf("%s != %s", v, jockId)
	}
}
