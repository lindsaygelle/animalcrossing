package animals

import "testing"

func TestFurSeal(t *testing.T) {
	if ok := FurSeal.Name() == furSeal; !ok {
		t.Fatal("FurlSeal.Name() != furSeal")
	}
}
