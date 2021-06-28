package personalities

import "testing"

func TestSisterlyName(t *testing.T) {
	if ok := Sisterly.Name() == sisterly; !ok {
		t.Fatal("Sisterly.Name() != sisterly")
	}
}
