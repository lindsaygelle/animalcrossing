package animals

import (
	"strings"
	"testing"
)

func testAnimals(animals []Animal, character string, t *testing.T) {
	for i := 0; i < len(animals); i++ {
		testAnimalsCharacter(animals[i], i, character, t)
		testAnimalsTitle(animals[i], i, t)
	}
}

func testAnimalsCharacter(animal Animal, i int, character string, t *testing.T) {
	var s string = animal.Name()
	var c string = string(s[0])
	if ok := c == character; !ok {
		t.Fatalf("%s != %s; i=%d", c, character, i)
	}
}

func testAnimalsTitle(animal Animal, i int, t *testing.T) {
	var s string = strings.Title(animal.Name())
	if ok := animal.Name() == s; !ok {
		t.Fatalf("%s != %s; i=%d", animal.Name(), s, i)
	}
}
