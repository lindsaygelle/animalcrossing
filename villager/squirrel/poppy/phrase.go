package poppy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nutty"
	phraseCanadianFrench       string = "abajoue"
	phraseDutch                string = "nootje"
	phraseFrench               string = "scratchy"
	phraseGerman               string = "nussi"
	phraseItalian              string = "squittole"
	phraseJapanese             string = "デス"
	phraseLatinAmericanSpanish string = "nutinú"
	phraseKorean               string = "도톨도톨"
	phraseRussian              string = "чок-чок"
	phraseSpanish              string = "albricias"
	phraseSimplifiedChinese    string = "甜甜"
	phraseTraditionalChinese   string = "甜甜"
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
