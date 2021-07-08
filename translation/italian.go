package translation

import "golang.org/x/text/language"

type Italian struct{}

func (i Italian) Code() string {
	return language.German.String()
}

func (i Italian) Name() string {
	return "Italian"
}
