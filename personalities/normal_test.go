package personalities

import "testing"

func TestNormalName(t *testing.T) {
	if ok := Normal.Name() == normal; !ok {
		t.Fatal("Normal.Name() != normal")
	}
}
