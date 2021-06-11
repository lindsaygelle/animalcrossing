package animalcrossing

import "testing"

func TestWalrusName(t *testing.T) {
    if ok := Walrus.Name() == walrus; !ok {
        t.Fatal("Walrus.Name() != walrus")
    }
}