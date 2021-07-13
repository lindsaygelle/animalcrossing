package frill_necked_lizard

type FrillNeckedLizard struct{}

func (f FrillNeckedLizard) Id() string    { return "frill_necked_lizard" }
func (f FrillNeckedLizard) Special() bool { return true }
