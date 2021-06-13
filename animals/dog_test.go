package animals

import "testing"

func TestDogName(t *testing.T) {
	if ok := Dog.Name() == dog; !ok {
		t.Fatal("Dog.Name() != dog")
	}
}
