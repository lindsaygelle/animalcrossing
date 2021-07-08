package translation

import "golang.org/x/text/language"

type English struct{}

func (e English) Code() string {
	return language.English.String()
}

func (e English) Name() string {
	return "English"
}
