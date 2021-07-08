package translations

import (
	"golang.org/x/text/language"
)

type Dutch struct{}

func (d Dutch) Code() string {
	return language.Dutch.String()
}
func (d Dutch) Name() string {
	return "Dutch"
}
