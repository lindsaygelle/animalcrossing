package friga

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "brrrmph"
	phraseCanadianFrench       string = "bourfff"
	phraseDutch                string = "brrrhmpf"
	phraseFrench               string = "bourfff"
	phraseGerman               string = "grrmpf"
	phraseItalian              string = "brrrumf"
	phraseJapanese             string = "ツルルン"
	phraseLatinAmericanSpanish string = "urfff"
	phraseKorean               string = "으쓱"
	phraseRussian              string = "брр-хм"
	phraseSpanish              string = "heladito"
	phraseSimplifiedChinese    string = "滑溜溜"
	phraseTraditionalChinese   string = "滑溜溜"
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
