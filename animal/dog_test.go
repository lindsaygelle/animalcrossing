package animal

import "testing"

// TestDogId tests whether Dog has the correct ID.
func TestDogId(t *testing.T) {
	if v := Dog.Id(); !(v == dogId) {
		t.Fatalf("%s != %s", v, dogId)
	}	
}

// TestDogFictional test whether Dog is a fictional Animal Crossing animal type.
func TestDogFictional(t *testing.T) {
	if v := Dog.Fictional(); !(v == dogFictional) {
		t.Fatalf("%t != %t", v, dogFictional)
	}
}

// TestDogSpecial tests whether Dog is a special Animal Crossing animal type.
func TestDogSpecial(t *testing.T) {
	if v := Dog.Special(); !(v == dogSpecial) {
		t.Fatalf("%t != %t", v, dogSpecial)
	}
}
