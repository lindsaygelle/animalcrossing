package monkey

import (
	"github.com/lindsaygelle/animalcrossing/animals/monkey/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type Monkey struct{}

<<<<<<< HEAD
func (m Monkey) Id() string {
	return "monkey"
}

func (m Monkey) Special() bool {
	return false
}

=======
func (m Monkey) Id() string    { return "monkey" }
func (m Monkey) Special() bool { return false }
>>>>>>> 27f353edaecf25fa7f87a238c128539f92f5550c
func (m Monkey) Translations() []languages.Translation {
	return []languages.Translation{
		translations.De{},
		translations.En{},
		translations.EsFemine{},
		translations.Es{},
		translations.FrFemine{},
		translations.Fr{},
		translations.It{},
		translations.Jp{},
		translations.Ko{},
		translations.Nl{},
		translations.Ru{},
		translations.Zh{}}
}
