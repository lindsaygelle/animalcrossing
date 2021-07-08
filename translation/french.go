package translation

import "golang.org/x/text/language"

type French struct{}

func (f French) Code() string {
	return language.French.String()
}

func (f French) Name() string {
	return "French"
}
