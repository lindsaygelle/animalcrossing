package translations

import "golang.org/x/text/language"

type Korean struct{}

func (k Korean) Code() string {
	return language.Korean.String()
}

func (k Korean) Name() string {
	return "Korean"
}
