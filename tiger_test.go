package animalcrossing

import "testing"

func TestTigerName(t *testing.T) {
    if ok := Tiger.Name() == tiger; !ok {
        t.Fatal("Tiger.Name() != tiger")
    }
}