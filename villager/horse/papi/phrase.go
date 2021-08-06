package papi

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "haaay"
	phraseCanadianFrench       string = "en selle"
	phraseDutch                string = "hooooi"
	phraseFrench               string = "en selle"
	phraseGerman               string = "nichja"
	phraseItalian              string = "clip clop"
	phraseJapanese             string = "だっけ"
	phraseLatinAmericanSpanish string = "cli-cló"
	phraseKorean               string = "그런가"
	phraseRussian              string = "се-е-ено"
	phraseSpanish              string = "quepassa"
	phraseSimplifiedChinese    string = "听说是"
	phraseTraditionalChinese   string = "聽說是"
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
