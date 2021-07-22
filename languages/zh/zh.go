package zh

import "github.com/lindsaygelle/animalcrossing/genders"

type Zh struct {
	genders.None
}

func (z Zh) Code() string {
	return "zh"
}

func (z Zh) Local() string {
	return "中文"
}

func (z Zh) Name() string {
	return "Chinese"
}
