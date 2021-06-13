package animals

import "testing"

func TestGiraffeName(t *testing.T) {
	if ok := Giraffe.Name() == giraffe; !ok {
		t.Fatal("Giraffe.Name() != giraffe")
	}
}
