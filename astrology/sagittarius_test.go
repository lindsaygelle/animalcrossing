package astrology

import "testing"

func TestSagittariusName(t *testing.T) {
	if ok := Sagittarius.Name() == sagittarius; !ok {
		t.Fatalf("%s != %s", Sagittarius.Name(), sagittarius)
	}
}
