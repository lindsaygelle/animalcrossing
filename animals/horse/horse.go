package horse

type Horse struct{}

func (h Horse) Id() string    { return "horse" }
func (h Horse) Special() bool { return false }
