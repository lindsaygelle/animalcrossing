package personality

import "testing"

// TestNormalId tests whether Normal has the correct ID.
func TestNormalId(t *testing.T) {
	if v := Normal.Id(); !(v == normalId) {
		t.Fatalf("%s != %s", v, normalId)
	}
}
