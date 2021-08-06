package peanut

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "slacker"
	phraseCanadianFrench       string = "noisette"
	phraseDutch                string = "luilak"
	phraseFrench               string = "noisette"
	phraseGerman               string = "faulenzer"
	phraseItalian              string = "sgronch"
	phraseJapanese             string = "なのよ"
	phraseLatinAmericanSpanish string = "esñik"
	phraseKorean               string = "거얌"
	phraseRussian              string = "лень"
	phraseSpanish              string = "almendrita"
	phraseSimplifiedChinese    string = "就是唷"
	phraseTraditionalChinese   string = "就是唷"
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
