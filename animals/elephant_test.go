package animals

import "testing"

func TestElephantName(t *testing.T) {
	if ok := Elephant.Name() == elephant; !ok {
		t.Fatal("Elephant.Name() != elephant")
	}
}
