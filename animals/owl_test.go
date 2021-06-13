package animals

import "testing"

func TestOwlName(t *testing.T) {
	if ok := Owl.Name() == owl; !ok {
		t.Fatal("Owl.Name() != owl")
	}
}
