package animal

import "testing"

// TestRabbitId tests whether Rabbit has the correct ID.
func TestRabbitId(t *testing.T) {
	if v := Rabbit.Id(); !(v == rabbitId) {
		t.Fatalf("%s != %s", v, rabbitId)
	}	
}

// TestRabbitFictional test whether Rabbit is a fictional Animal Crossing animal type.
func TestRabbitFictional(t *testing.T) {
	if v := Rabbit.Fictional(); !(v == rabbitFictional) {
		t.Fatalf("%t != %t", v, rabbitFictional)
	}
}

// TestRabbitSpecial tests whether Rabbit is a special Animal Crossing animal type.
func TestRabbitSpecial(t *testing.T) {
	if v := Rabbit.Special(); !(v == rabbitSpecial) {
		t.Fatalf("%t != %t", v, rabbitSpecial)
	}
}
