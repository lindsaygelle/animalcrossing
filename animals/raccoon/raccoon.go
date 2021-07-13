package raccoon

type Raccoon struct{}

func (r Raccoon) Id() string    { return "raccoon" }
func (r Raccoon) Special() bool { return true }
