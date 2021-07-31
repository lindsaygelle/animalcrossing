package astrology

// Astrology is an astrologial star sign for an Animal Crossing villager.
type Astrology interface {
	Id() string
	Longtitude() uint16
	Names() Names
	Symbol() string
	Unicode() (string, bool)
}

// astrology implements Astrology.
type astrology struct {
	id         string
	longtitude uint16
	names      names
	symbol     string
	unicode    string
}

func (v astrology) Id() string {
	return v.id
}

func (v astrology) Longtitude() uint16 {
	return v.longtitude
}

func (v astrology) Names() Names {
	return v.names
}

func (v astrology) Symbol() string {
	return v.symbol
}

func (v astrology) Unicode() (string, bool) {
	return v.unicode, (len(v.unicode) > 0)
}

var (
	_ Astrology = astrology{}
)
