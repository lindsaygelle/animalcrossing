package chai

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "flap flap"
	phraseCanadianFrench       string = "pachipach"
	phraseDutch                string = "kaneel"
	phraseFrench               string = "pachipach"
	phraseGerman               string = "zimtroll"
	phraseItalian              string = "ruuu ruuu"
	phraseJapanese             string = "ロルロル"
	phraseLatinAmericanSpanish string = "vuelalto"
	phraseKorean               string = "롤롤"
	phraseRussian              string = "машу ушами"
	phraseSpanish              string = "vuelalto"
	phraseSimplifiedChinese    string = "大耳"
	phraseTraditionalChinese   string = "大耳"
)

var (
	phrase = map[language.Tag]string{
		language.AmericanEnglish:      phraseAmericanEnglish,
		language.CanadianFrench:       phraseCanadianFrench,
		language.Dutch:                phraseDutch,
		language.French:               phraseFrench,
		language.German:               phraseGerman,
		language.Italian:              phraseItalian,
		language.Japanese:             phraseJapanese,
		language.Korean:               phraseKorean,
		language.LatinAmericanSpanish: phraseLatinAmericanSpanish,
		language.Russian:              phraseRussian,
		language.Spanish:              phraseSpanish,
		language.SimplifiedChinese:    phraseSimplifiedChinese,
		language.TraditionalChinese:   phraseTraditionalChinese}
)
