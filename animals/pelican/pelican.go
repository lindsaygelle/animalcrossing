package pelican

type Pelican struct{}

func (p Pelican) Id() string    { return "pelican" }
func (p Pelican) Special() bool { return true }
