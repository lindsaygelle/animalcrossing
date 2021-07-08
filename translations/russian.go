package translations

import "golang.org/x/text/language"

type Russian struct{}

func (r Russian) Code() string {
	return language.Russian.String()
}

func (r Russian) Name() string {
	return "Russian"
}
