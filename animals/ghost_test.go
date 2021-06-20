package animals

import "testing"

func TestGhostName(t *testing.T) {
	if ok := Ghost.Name() == ghost; !ok {
		t.Fatal("Ghost.Name() != ghost")
	}
}
