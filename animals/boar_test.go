package animals

import "testing"

func TestBoarName(t *testing.T) {
	if ok := Boar.Name() == boar; !ok {
		t.Fatal("Boar.Name() != boar")
	}
}
