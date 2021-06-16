package astrology

// Astrology is an astrological star sign for an Animal Crossing villager.
type Astrology interface {
	// Name is the name of the astrological star sign.
	Name() string
}

// astrology implements the Astrology interface.
type astrology struct {
	name string
}

// Name returns the astrologys name.
func (a astrology) Name() string {
	return a.name
}

var (
	// validate astrology implements Astrology.
	_ Astrology = astrology{}
)

var (
	All = [...]Astrology{
		Aquarius,
		Aries,
		Cancer,
		Capricorn,
		Gemini,
		Leo,
		Libra,
		Pisces,
		Sagittarius,
		Scorpio,
		Taurus,
		Virgo}
)
