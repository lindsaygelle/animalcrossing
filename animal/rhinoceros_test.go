package animal

import "testing"

// TestRhinocerosId tests whether Rhinoceros has the correct ID.
func TestRhinocerosId(t *testing.T) {
	if v := Rhinoceros.Id(); !(v == rhinocerosId) {
		t.Fatalf("%s != %s", v, rhinocerosId)
	}	
}

// TestRhinocerosFictional test whether Rhinoceros is a fictional Animal Crossing animal type.
func TestRhinocerosFictional(t *testing.T) {
	if v := Rhinoceros.Fictional(); !(v == rhinocerosFictional) {
		t.Fatalf("%t != %t", v, rhinocerosFictional)
	}
}

// TestRhinocerosSpecial tests whether Rhinoceros is a special Animal Crossing animal type.
func TestRhinocerosSpecial(t *testing.T) {
	if v := Rhinoceros.Special(); !(v == rhinocerosSpecial) {
		t.Fatalf("%t != %t", v, rhinocerosSpecial)
	}
}
