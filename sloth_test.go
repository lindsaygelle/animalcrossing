package animalcrossing

import "testing"

func TestSlothName(t *testing.T) {
    if ok := Sloth.Name() == sloth; !ok {
        t.Fatal("Sloth.Name() != sloth")
    }
}
