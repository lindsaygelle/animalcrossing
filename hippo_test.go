package animalcrossing

import "testing"

func TestHippoName(t *testing.T) {
	if ok := Hippo.Name() == hippo; !ok {
		t.Fatal("Hippo.Name() != hippo")
	}
}
