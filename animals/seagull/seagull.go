package seagull

type Seagull struct{}

func (s Seagull) Id() string    { return "seagull" }
func (s Seagull) Special() bool { return true }
