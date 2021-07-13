package axolotl

type Axolotl struct{}

func (a Axolotl) Id() string    { return "axolotl" }
func (a Axolotl) Special() bool { return true }
