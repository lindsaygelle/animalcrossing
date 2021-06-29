package astrology

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAstrologyName(t *testing.T) {
	var astrologyName string = fmt.Sprint(rand.Int())
	var astrology Astrology = astrology{
		name: astrologyName}
	if ok := astrology.Name() == astrologyName; !ok {
		t.Fatalf("%s != %s", astrology.Name(), astrologyName)
	}
}
