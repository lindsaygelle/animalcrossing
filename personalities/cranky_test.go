package personalities

import "testing"

func TestCrankyName(t *testing.T) {
	if ok := Cranky.Name() == cranky; !ok {
		t.Fatal("Cranky.Name() != cranky")
	}
}
