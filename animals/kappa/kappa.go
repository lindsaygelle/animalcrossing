package kappa

type Kappa struct{}

func (k Kappa) Id() string    { return "kappa" }
func (k Kappa) Special() bool { return true }
