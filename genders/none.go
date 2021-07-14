package genders

type None struct{}

func (n None) Gender() string {
	return "None"
}
