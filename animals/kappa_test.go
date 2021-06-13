package animals

import "testing"

func TestKappaName(t *testing.T) {
	if ok := Kappa.Name() == kappa; !ok {
		t.Fatalf("Kappa.Name() != kappa")
	}
}
