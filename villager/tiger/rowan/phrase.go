package rowan

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mango"
	phraseCanadianFrench       string = "banane"
	phraseDutch                string = "mango"
	phraseFrench               string = "banane"
	phraseGerman               string = "mango"
	phraseItalian              string = "mango"
	phraseJapanese             string = "まったく"
	phraseLatinAmericanSpanish string = "gurruf"
	phraseKorean               string = "내참"
	phraseRussian              string = "манго"
	phraseSpanish              string = "turista"
	phraseSimplifiedChinese    string = "完全"
	phraseTraditionalChinese   string = "完全"
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
