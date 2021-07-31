package astrology

import "golang.org/x/text/language"

// Astrologys is a collection of astrology.Astrology.
type Astrologys interface {
	ForEach(func(language.Tag, Astrology))
	Get(language.Tag) (Astrology, bool)
	Len() int
}

// astrologys implements Astrologys.
type astrologys map[language.Tag]astrology

func (v astrologys) ForEach(fn func(language.Tag, Astrology)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v astrologys) Get(k language.Tag) (Astrology, bool) {
	astrology, ok := v[k]
	return astrology, ok
}

func (v astrologys) Len() int {
	return len(v)
}

var (
	_ Astrologys = astrologys{}
)
