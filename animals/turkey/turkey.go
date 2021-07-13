package turkey

type Turkey struct{}

func (t Turkey) Id() string    { return "turkey" }
func (t Turkey) Special() bool { return true }
