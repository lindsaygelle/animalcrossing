package animalcrossing

import "testing"

func TestRaccoonName(t *testing.T) {
    if ok := Raccoon.Name() == raccoon; !ok {
        t.Fatal("Raccoon.Name() != raccoon")
    }
}
