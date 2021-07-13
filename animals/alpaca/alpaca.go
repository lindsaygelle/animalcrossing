package alpaca

type Alpaca struct{}

func (a Alpaca) Id() string    { return "alpaca" }
func (a Alpaca) Special() bool { return true }
