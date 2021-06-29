package genders

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGenderName(t *testing.T) {
	var genderName string = fmt.Sprint(rand.Int())
	var gender Gender = gender{
		name: genderName}
	if ok := gender.Name() == genderName; !ok {
		t.Fatalf("%s != %s", gender.Name(), genderName)
	}
}
