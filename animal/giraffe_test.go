package animal

import "testing"

// TestGiraffeId tests whether Giraffe has the correct ID.
func TestGiraffeId(t *testing.T) {
	if v := Giraffe.Id(); !(v == giraffeId) {
		t.Fatalf("%s != %s", v, giraffeId)
	}	
}

// TestGiraffeFictional test whether Giraffe is a fictional Animal Crossing animal type.
func TestGiraffeFictional(t *testing.T) {
	if v := Giraffe.Fictional(); !(v == giraffeFictional) {
		t.Fatalf("%t != %t", v, giraffeFictional)
	}
}

// TestGiraffeSpecial tests whether Giraffe is a special Animal Crossing animal type.
func TestGiraffeSpecial(t *testing.T) {
	if v := Giraffe.Special(); !(v == giraffeSpecial) {
		t.Fatalf("%t != %t", v, giraffeSpecial)
	}
}
