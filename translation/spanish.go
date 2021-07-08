package translation

import "golang.org/x/text/language"

type Spanish struct{}

func (s Spanish) Code() string {
	return language.Spanish.String()
}

func (s Spanish) Name() string {
	return "Spanish"
}
