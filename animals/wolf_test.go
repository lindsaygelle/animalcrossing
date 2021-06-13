package animals

import "testing"

func TestWolfName(t *testing.T) {
	if ok := Wolf.Name() == wolf; !ok {
		t.Fatal("Wolf.Name() != wolf")
	}
}
