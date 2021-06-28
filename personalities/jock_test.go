package personalities

import "testing"

func TestJockName(t *testing.T) {
	if ok := Jock.Name() == jock; !ok {
		t.Fatal("Jock.Name() != jock")
	}
}
