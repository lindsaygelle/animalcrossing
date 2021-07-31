package animal

import "testing"

// TestFrillneckedLizardId tests whether FrillneckedLizard has the correct ID.
func TestFrillneckedLizardId(t *testing.T) {
	if v := FrillneckedLizard.Id(); !(v == frillneckedlizardId) {
		t.Fatalf("%s != %s", v, frillneckedlizardId)
	}	
}

// TestFrillneckedLizardFictional test whether FrillneckedLizard is a fictional Animal Crossing animal type.
func TestFrillneckedLizardFictional(t *testing.T) {
	if v := FrillneckedLizard.Fictional(); !(v == frillneckedlizardFictional) {
		t.Fatalf("%t != %t", v, frillneckedlizardFictional)
	}
}

// TestFrillneckedLizardSpecial tests whether FrillneckedLizard is a special Animal Crossing animal type.
func TestFrillneckedLizardSpecial(t *testing.T) {
	if v := FrillneckedLizard.Special(); !(v == frillneckedlizardSpecial) {
		t.Fatalf("%t != %t", v, frillneckedlizardSpecial)
	}
}
