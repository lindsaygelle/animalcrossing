package jp

import "github.com/lindsaygelle/animalcrossing/genders"

type Jp struct {
	genders.None
}

func (j Jp) Code() string {
	return "jp"
}

func (j Jp) Local() string {
	return "日本語"
}

func (j Jp) Name() string {
	return "Japanese"
}
