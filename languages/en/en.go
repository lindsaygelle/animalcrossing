package en

import "github.com/lindsaygelle/animalcrossing/genders"

type En struct {
	genders.None
}

func (e En) Code() string {
	return "en"
}

func (e En) Local() string {
	return "English"
}

func (e En) Name() string {
	return e.Local()
}
