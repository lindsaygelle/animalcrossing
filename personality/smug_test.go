package personality

import "testing"

// TestSmugId tests whether Smug has the correct ID.
func TestSmugId(t *testing.T) {
	if v := Smug.Id(); !(v == smugId) {
		t.Fatalf("%s != %s", v, smugId)
	}
}
