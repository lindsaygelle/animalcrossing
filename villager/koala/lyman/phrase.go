package lyman

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chips"
	phraseCanadianFrench       string = "sandwich"
	phraseDutch                string = "chips"
	phraseFrench               string = "sandwich"
	phraseGerman               string = "liptus"
	phraseItalian              string = "ponzipò"
	phraseJapanese             string = "わ～い"
	phraseLatinAmericanSpanish string = "tralalá"
	phraseKorean               string = "우와"
	phraseRussian              string = "ничего"
	phraseSpanish              string = "trololó"
	phraseSimplifiedChinese    string = "哇"
	phraseTraditionalChinese   string = "哇～"
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
