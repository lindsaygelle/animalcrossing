package ricky

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nutcase"
	phraseCanadianFrench       string = "dingo"
	phraseDutch                string = "mafnoot"
	phraseFrench               string = "dingo"
	phraseGerman               string = "knorz"
	phraseItalian              string = "nocciola"
	phraseJapanese             string = "キッ"
	phraseLatinAmericanSpanish string = "crac-crac"
	phraseKorean               string = "사각사각"
	phraseRussian              string = "арахис"
	phraseSpanish              string = "cáscaras"
	phraseSimplifiedChinese    string = "炯炯"
	phraseTraditionalChinese   string = "炯炯"
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
