package fur_seal

type FurSeal struct{}

func (f FurSeal) Id() string    { return "fur_seal" }
func (f FurSeal) Special() bool { return true }
