package huck

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hopper"
	phraseCanadianFrench       string = "nouille"
	phraseDutch                string = "hophop"
	phraseFrench               string = "nouille"
	phraseGerman               string = "hops"
	phraseItalian              string = "hophop"
	phraseJapanese             string = "とかー"
	phraseLatinAmericanSpanish string = "ajop"
	phraseKorean               string = "싸우자"
	phraseRussian              string = "скок"
	phraseSpanish              string = "grillito"
	phraseSimplifiedChinese    string = "喝喝"
	phraseTraditionalChinese   string = "喝喝"
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
