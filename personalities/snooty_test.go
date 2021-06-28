package personalities

import "testing"

func TestSnootyName(t *testing.T) {
	if ok := Snooty.Name() == snooty; !ok {
		t.Fatal("Snooty.Name() != snooty")
	}
}
