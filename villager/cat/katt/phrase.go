package katt

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "purrty"
	phraseCanadianFrench       string = "miaouille"
	phraseDutch                string = "prrrachtig"
	phraseFrench               string = "miaouille"
	phraseGerman               string = "zischel"
	phraseItalian              string = "fffz"
	phraseJapanese             string = "じゃね"
	phraseLatinAmericanSpanish string = "purrrs"
	phraseKorean               string = "휘리릭"
	phraseRussian              string = "хор-р-рошо"
	phraseSpanish              string = "salmonete"
	phraseSimplifiedChinese    string = "掰掰"
	phraseTraditionalChinese   string = "掰掰"
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
