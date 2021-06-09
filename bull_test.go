package animalcrossing

import "testing"

func TestBullName(t *testing.T) {
	if ok := Bull.Name() == bull; !ok {
		t.Fatal("Bull.Name() != bull")
	}
}
