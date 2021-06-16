package astrology

import "testing"

func TestAquariusName(t *testing.T) {
	if ok := Aquarius.Name() == aquarius; !ok {
		t.Fatalf("%s != %s", Aquarius.Name(), aquarius)
	}
}
