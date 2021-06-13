package animals

import "testing"

func TestTapirName(t *testing.T) {
	if ok := Tapir.Name() == tapir; !ok {
		t.Fatal("Tapir.Name() != tapir")
	}
}
