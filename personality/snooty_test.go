package personality

import "testing"

// TestSnootyId tests whether Snooty has the correct ID.
func TestSnootyId(t *testing.T) {
	if v := Snooty.Id(); !(v == snootyId) {
		t.Fatalf("%s != %s", v, snootyId)
	}
}
