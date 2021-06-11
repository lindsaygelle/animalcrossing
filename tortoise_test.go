package animalcrossing

import "testing"

func TestTortoiseName(t *testing.T) {
    if ok := Tortoise.Name() == tortoise; !ok {
        t.Fatal("Tortoise.Name() != tortoise")
    }
}




