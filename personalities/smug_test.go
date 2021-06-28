package personalities

import "testing"

func TestSmugName(t *testing.T) {
	if ok := Smug.Name() == smug; !ok {
		t.Fatal("Smug.Name() != smug")
	}
}
