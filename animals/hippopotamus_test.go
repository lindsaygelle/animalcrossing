package animals

import "testing"

func TestHippopotamusName(t *testing.T) {
	if ok := Hippopotamus.Name() == hippopotamus; !ok {
		t.Fatal("Hippopotamus.Name() != hippopotamus")
	}
}
