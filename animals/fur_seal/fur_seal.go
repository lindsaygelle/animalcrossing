package fur_seal

import (
	"github.com/lindsaygelle/animalcrossing/animals/fur_seal/translations"
	"github.com/lindsaygelle/animalcrossing/languages"
)

type FurSeal struct{}

func (f FurSeal) Id() string {
	return "fur_seal"
}

func (f FurSeal) Special() bool {
	return true
}

func (f FurSeal) Translations() []languages.Translation {
	return []languages.Translation{
		translations.En{}}
}
