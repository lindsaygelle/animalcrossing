package animalcrossing

import "testing"

func TestFrillNeckedLizard(t *testing.T) {
	if ok := FrillNeckedLizard.Name() == frillNeckedLizard; !ok {
		t.Fatal("FrillNeckedLizard.Name() != frillNeckedLizard")
	}
}
