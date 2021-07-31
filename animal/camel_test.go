package animal

import "testing"

// TestCamelId tests whether Camel has the correct ID.
func TestCamelId(t *testing.T) {
	if v := Camel.Id(); !(v == camelId) {
		t.Fatalf("%s != %s", v, camelId)
	}	
}

// TestCamelFictional test whether Camel is a fictional Animal Crossing animal type.
func TestCamelFictional(t *testing.T) {
	if v := Camel.Fictional(); !(v == camelFictional) {
		t.Fatalf("%t != %t", v, camelFictional)
	}
}

// TestCamelSpecial tests whether Camel is a special Animal Crossing animal type.
func TestCamelSpecial(t *testing.T) {
	if v := Camel.Special(); !(v == camelSpecial) {
		t.Fatalf("%t != %t", v, camelSpecial)
	}
}
