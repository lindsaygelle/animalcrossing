package genders

// Gender is the gender of an Animal Crossing villager.
type Gender interface {
	// Name is the name of the gender.
	Name() string
}

// gender implements the Gender interface.
type gender struct {
	name string
}

// Name returns the genders name.
func (g gender) Name() string {
	return g.name
}

var (
	// validate gender implements Gender.
	_ Gender = gender{}
)
