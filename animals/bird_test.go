package animals

import "testing"

func TestBirdName(t *testing.T) {
	if ok := Bird.Name() == bird; !ok {
		t.Fatal("Bird.Name() != bird")
	}
}
