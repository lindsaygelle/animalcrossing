package animal

import "testing"

// TestCatId tests whether Cat has the correct ID.
func TestCatId(t *testing.T) {
	if v := Cat.Id(); !(v == catId) {
		t.Fatalf("%s != %s", v, catId)
	}	
}

// TestCatFictional test whether Cat is a fictional Animal Crossing animal type.
func TestCatFictional(t *testing.T) {
	if v := Cat.Fictional(); !(v == catFictional) {
		t.Fatalf("%t != %t", v, catFictional)
	}
}

// TestCatSpecial tests whether Cat is a special Animal Crossing animal type.
func TestCatSpecial(t *testing.T) {
	if v := Cat.Special(); !(v == catSpecial) {
		t.Fatalf("%t != %t", v, catSpecial)
	}
}
