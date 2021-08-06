package tucker

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "fuzzers"
	phraseCanadianFrench       string = "mammouth"
	phraseDutch                string = "pluisie"
	phraseFrench               string = "mammouth"
	phraseGerman               string = "puuuuuh"
	phraseItalian              string = "barrir"
	phraseJapanese             string = "もじゃ"
	phraseLatinAmericanSpanish string = "purupú"
	phraseKorean               string = "뿌뿌"
	phraseRussian              string = "бивень"
	phraseSpanish              string = "purupú"
	phraseSimplifiedChinese    string = "毛毛"
	phraseTraditionalChinese   string = "毛毛"
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
