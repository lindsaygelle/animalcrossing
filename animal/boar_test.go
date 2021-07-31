package animal

import "testing"

// TestBoarId tests whether Boar has the correct ID.
func TestBoarId(t *testing.T) {
	if v := Boar.Id(); !(v == boarId) {
		t.Fatalf("%s != %s", v, boarId)
	}	
}

// TestBoarFictional test whether Boar is a fictional Animal Crossing animal type.
func TestBoarFictional(t *testing.T) {
	if v := Boar.Fictional(); !(v == boarFictional) {
		t.Fatalf("%t != %t", v, boarFictional)
	}
}

// TestBoarSpecial tests whether Boar is a special Animal Crossing animal type.
func TestBoarSpecial(t *testing.T) {
	if v := Boar.Special(); !(v == boarSpecial) {
		t.Fatalf("%t != %t", v, boarSpecial)
	}
}
