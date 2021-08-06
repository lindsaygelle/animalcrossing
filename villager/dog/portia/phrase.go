package portia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ruffian"
	phraseCanadianFrench       string = "brute"
	phraseDutch                string = "keffer"
	phraseFrench               string = "brute"
	phraseGerman               string = "frechi"
	phraseItalian              string = "ronfietto"
	phraseJapanese             string = "フンッ"
	phraseLatinAmericanSpanish string = "ruf-ruf"
	phraseKorean               string = "흥"
	phraseRussian              string = "тяв-ряв"
	phraseSpanish              string = "ruf-ruf"
	phraseSimplifiedChinese    string = "屋屋"
	phraseTraditionalChinese   string = "屋屋"
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
