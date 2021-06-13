package animals

import "testing"

func TestGullName(t *testing.T) {
	if ok := Gull.Name() == gull; !ok {
		t.Fatal("Gull.Name() != gull")
	}
}
