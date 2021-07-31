package translation

import (
	"golang.org/x/text/language"
)

// Translations is a collection of Translation.
type Translations interface {
	ForEach(func(language.Tag, Translation))
	Get(language.Tag) (Translation, bool)
	Len() int
	Must(language.Tag) Translation
}

// translations implements Translations.
type translations map[language.Tag]Translation

func (v translations) Add(t Translation) {
	v[t.Language()] = t
}

func (v translations) ForEach(fn func(language.Tag, Translation)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v translations) Get(k language.Tag) (Translation, bool) {
	translation, ok := v[k]
	return translation, ok
}

func (v translations) Len() int {
	return len(v)
}

func (v translations) Must(k language.Tag) Translation {
	translation, ok := v.Get(k)
	if !ok {
		panic(k)
	}
	return translation
}

var (
	_ Translations = translations{}
)

// NewTranslation returns a new Translations.
func NewTranslations(t ...Translation) Translations {
	translations := translations{}
	for _, v := range t {
		translations.Add(v)
	}
	return translations
}
