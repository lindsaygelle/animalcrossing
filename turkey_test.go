package animalcrossing

import "testing"

func TestTurkeyName(t *testing.T) {
    if ok := Turkey.Name() == turkey; !ok {
        t.Fatal("Turkey.Name() != turkey")
    }
}