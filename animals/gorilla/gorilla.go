package gorilla

type Gorilla struct{}

func (g Gorilla) Id() string    { return "gorilla" }
func (g Gorilla) Special() bool { return false }
