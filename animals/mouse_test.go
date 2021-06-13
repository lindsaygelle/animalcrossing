package animals

import "testing"

func TestMouseName(t *testing.T) {
	if ok := Mouse.Name() == mouse; !ok {
		t.Fatal("Mouse.Name() != mouse")
	}
}
