package species

import (
	"reflect"
	"testing"
)

func TestSpecies(t *testing.T) {
	var s Species = species{}
	if ok := reflect.ValueOf(s).IsZero(); !ok {
		t.Fatal("species{} is zero value")
	}
}
