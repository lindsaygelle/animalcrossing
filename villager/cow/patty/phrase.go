package patty

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "how-now"
	phraseCanadianFrench       string = "oui oui"
	phraseDutch                string = "koebeest"
	phraseFrench               string = "allez"
	phraseGerman               string = "muuhna"
	phraseItalian              string = "truuulà"
	phraseJapanese             string = "だモー"
	phraseLatinAmericanSpanish string = "muuuu"
	phraseKorean               string = "음메"
	phraseRussian              string = "му-му"
	phraseSpanish              string = "muuuu"
	phraseSimplifiedChinese    string = "牟"
	phraseTraditionalChinese   string = "牟"
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
