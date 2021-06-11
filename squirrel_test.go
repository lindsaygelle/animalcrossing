package animalcrossing

import "testing"

func TestSquirrelName(t *testing.T) {
    if ok := Squirrel.Name() == squirrel; !ok {
        t.Fatal("Squirrel.Name() != squirrel")
    }
}