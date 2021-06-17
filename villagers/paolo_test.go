
package villagers 

import (
    "testing"

    "github.com/lindsaygelle/animalcrossing/animals"
)

func TestPaoloAnimal(t *testing.T) {
    var s string = animals.Elephant.Name()
    if ok := Paolo.Animal() == s; !ok {
        t.Fatalf("%s != %s", Paolo.Animal(), s)
    }
}

func TestPaoloName(t *testing.T) {
    if ok := Paolo.Name() == paolo; !ok {
        t.Fatalf("%s != %s", Paolo.Name(), paolo)
    }
}