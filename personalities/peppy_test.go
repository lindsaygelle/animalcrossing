package personalities

import "testing"

func TestPeppyName(t *testing.T) {
	if ok := Peppy.Name() == peppy; !ok {
		t.Fatal("Peppy.Name() != peppy")
	}
}
