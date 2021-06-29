package animals

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAnimalName(t *testing.T) {
	var animalName string = fmt.Sprint(rand.Int())
	var animal Animal = animal{
		name: animalName}
	if ok := animal.Name() == animalName; !ok {
		t.Fatalf("%s != %s", animal.Name(), animalName)
	}
}
