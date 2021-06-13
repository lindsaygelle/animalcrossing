package animals

import "testing"

func TestRhinocerosName(t *testing.T) {
	if ok := Rhinoceros.Name() == rhinoceros; !ok {
		t.Fatal("Rhinoceros.Name() != rhinoceros")
	}
}
