package personality

import "testing"

// TestLazyId tests whether Lazy has the correct ID.
func TestLazyId(t *testing.T) {
	if v := Lazy.Id(); !(v == lazyId) {
		t.Fatalf("%s != %s", v, lazyId)
	}
}
