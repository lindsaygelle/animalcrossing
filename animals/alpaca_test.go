package animals

import "testing"

func TestAlpacaName(t *testing.T) {
	if ok := Alpaca.Name() == alpaca; !ok {
		t.Fatal("Alpaca.Name() != alpaca")
	}
}
