package villager

import "golang.org/x/text/language"

// Catchphrases is a collection of Animal Crossing villager catchphrase in various languages.
type Catchphrases interface {
	ForEach(fn func(language.Tag, Catchphrase))
	Get(language.Tag) (Catchphrase, bool)
	Len() int
}

// catchphrases implements Catchphrases.
type catchphrases map[language.Tag]catchphrase

func (v catchphrases) ForEach(fn func(language.Tag, Catchphrase)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v catchphrases) Get(k language.Tag) (Catchphrase, bool) {
	catchphrase, ok := v[k]
	return catchphrase, ok
}

func (v catchphrases) Len() int {
	return len(v)
}

var (
	_ Catchphrases = catchphrases{}
)
