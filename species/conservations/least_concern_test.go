package conservations

import "testing"

func TestLeastConcern(t *testing.T) {
	var s string = "Least Concern"
	if ok := leastConcern == s; !ok {
		t.Fatalf("leastConcern != %s", s)
	}
}
