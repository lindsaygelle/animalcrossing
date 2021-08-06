package peaches

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "neighbor"
	phraseCanadianFrench       string = "mustang"
	phraseDutch                string = "hinnik"
	phraseFrench               string = "mustang"
	phraseGerman               string = "nachbar"
	phraseItalian              string = "oh mamma"
	phraseJapanese             string = "だポン"
	phraseLatinAmericanSpanish string = "niiiih"
	phraseKorean               string = "몰라요"
	phraseRussian              string = "соседушка"
	phraseSpanish              string = "niiiih"
	phraseSimplifiedChinese    string = "蹦"
	phraseTraditionalChinese   string = "蹦"
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
