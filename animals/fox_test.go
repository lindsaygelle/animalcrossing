package animals

import "testing"

func TestFoxName(t *testing.T) {
	if ok := Fox.Name() == fox; !ok {
		t.Fatal("Fox.Name() != fox")
	}
}
