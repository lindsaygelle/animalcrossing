package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCookieAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Cookie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cookie.Animal(), s)
	}
}

func TestCookieName(t *testing.T) {
	if ok := Cookie.Name() == cookie; !ok {
		t.Fatalf("%s != %s", Cookie.Name(), cookie)
	}
}
