package animal

import "testing"

// TestTurtleId tests whether Turtle has the correct ID.
func TestTurtleId(t *testing.T) {
	if v := Turtle.Id(); !(v == turtleId) {
		t.Fatalf("%s != %s", v, turtleId)
	}	
}

// TestTurtleFictional test whether Turtle is a fictional Animal Crossing animal type.
func TestTurtleFictional(t *testing.T) {
	if v := Turtle.Fictional(); !(v == turtleFictional) {
		t.Fatalf("%t != %t", v, turtleFictional)
	}
}

// TestTurtleSpecial tests whether Turtle is a special Animal Crossing animal type.
func TestTurtleSpecial(t *testing.T) {
	if v := Turtle.Special(); !(v == turtleSpecial) {
		t.Fatalf("%t != %t", v, turtleSpecial)
	}
}
