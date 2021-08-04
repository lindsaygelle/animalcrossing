package animal

import "golang.org/x/text/language"

type Animal struct {
	Id   string
	Key  Key
	Name map[language.Tag]string
}
