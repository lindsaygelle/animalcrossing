package ketchup

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bitty"
	phraseCanadianFrench       string = "plouf"
	phraseDutch                string = "tomaatje"
	phraseFrench               string = "plouf"
	phraseGerman               string = "pullilol"
	phraseItalian              string = "quaraquack"
	phraseJapanese             string = "プチ"
	phraseLatinAmericanSpanish string = "cuic-cuic"
	phraseKorean               string = "찌익"
	phraseRussian              string = "покрякушка"
	phraseSpanish              string = "cuic-cuic"
	phraseSimplifiedChinese    string = "噗吱"
	phraseTraditionalChinese   string = "噗吱"
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
