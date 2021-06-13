package animals

import "testing"

func TestHorseName(t *testing.T) {
	if ok := Horse.Name() == horse; !ok {
		t.Fatal("Horse.Name() != horse")
	}
}
