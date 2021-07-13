package tortoise

type Tortoise struct{}

func (t Tortoise) Id() string    { return "tortoise" }
func (t Tortoise) Special() bool { return true }
