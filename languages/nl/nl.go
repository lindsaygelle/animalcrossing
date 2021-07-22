package nl

import "github.com/lindsaygelle/animalcrossing/genders"

type Nl struct {
	genders.None
}

func (n Nl) Code() string {
	return "nl"
}

func (n Nl) Local() string {
	return "Nederlands"
}

func (n Nl) Name() string {
	return "Dutch"
}
