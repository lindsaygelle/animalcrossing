package animalcrossing

import "testing"

func TestSkunkName(t *testing.T) {
    if ok := Skunk.Name() == skunk; !ok {
        t.Fatal("Skunk.Name() != skunk")
    }
}