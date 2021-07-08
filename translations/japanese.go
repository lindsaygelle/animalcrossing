package translations

import "golang.org/x/text/language"

type Japanese struct{}

func (j Japanese) Code() string {
	return language.Japanese.String()
}

func (j Japanese) Name() string {
	return "Japanese"
}
