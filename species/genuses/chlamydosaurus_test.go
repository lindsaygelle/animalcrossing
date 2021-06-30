package genuses

import "testing"

func TestChlamydosaurus(t *testing.T) {
	var s string = "Chlamydosaurus"
	if ok := chlamydosaurus == s; !ok {
		t.Fatalf("chlamydosaurus != %s", s)
	}
}
