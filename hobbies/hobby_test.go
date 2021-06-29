package hobbies

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHobbyName(t *testing.T) {
	var hobbyName string = fmt.Sprint(rand.Int())
	var hobby Hobby = hobby{
		name: hobbyName}
	if ok := hobby.Name() == hobbyName; !ok {
		t.Fatalf("%s != %s", hobby.Name(), hobbyName)
	}
}
