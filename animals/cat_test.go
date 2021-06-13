package animals

import "testing"

func TestCatName(t *testing.T) {
	if ok := Cat.Name() == cat; !ok {
		t.Fatal("Cat.Name() != cat")
	}
}
