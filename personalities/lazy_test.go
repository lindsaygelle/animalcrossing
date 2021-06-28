package personalities

import "testing"

func TestLazyName(t *testing.T) {
	if ok := Lazy.Name() == lazy; !ok {
		t.Fatal("Lazy.Name() != lazy")
	}
}
