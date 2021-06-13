package animals

import "testing"

func TestGyroidName(t *testing.T) {
	if ok := Gyroid.Name() == gyroid; !ok {
		t.Fatal("Gyroid.Name() != gyroid")
	}
}
