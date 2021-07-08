package translation

import "golang.org/x/text/language"

type German struct{}

func (g German) Code() string {
	return language.German.String()
}

func (g German) Name() string {
	return "German"
}
