package camel

type Camel struct{}

func (c Camel) Id() string    { return "camel" }
func (c Camel) Special() bool { return true }
