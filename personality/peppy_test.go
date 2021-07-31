package personality

import "testing"

// TestPeppyId tests whether Peppy has the correct ID.
func TestPeppyId(t *testing.T) {
	if v := Peppy.Id(); !(v == peppyId) {
		t.Fatalf("%s != %s", v, peppyId)
	}
}
