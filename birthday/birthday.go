package birthday

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

// Birthday is an Animal Crossing villagers birthday.
type Birthday interface {
	Day() uint8
	Month() time.Month
	Names() translation.Translations
}

// birthday implements Birthday.
type birthday struct {
	day   uint8
	month time.Month
	names translation.Translations
}

func (v birthday) Day() uint8 {
	return v.day
}

func (v birthday) Month() time.Month {
	return v.month
}

func (v birthday) Names() translation.Translations {
	return v.names
}

var (
	_ Birthday = birthday{}
)

// NewBirthday returns a new Birthday.
func NewBirthday(day uint8, month time.Month) Birthday {
	return birthday{
		day:   day,
		month: month,
		names: translation.NewTranslations(
			translation.NewTranslation(
				language.AmericanEnglish, month.String()))}
}
