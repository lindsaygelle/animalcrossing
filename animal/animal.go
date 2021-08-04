package animal

import "golang.org/x/text/language"

type Animal struct {
	Id   string
	Name map[language.Tag]string
}
