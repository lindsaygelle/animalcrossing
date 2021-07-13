package reindeer

type Reindeer struct{}

func (r Reindeer) Id() string    { return "reindeer" }
func (r Reindeer) Special() bool { return true }
