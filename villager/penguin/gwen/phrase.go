package gwen

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "h-h-h-hon"
	phraseCanadianFrench       string = "h-h-h-hon"
	phraseDutch                string = "d-d-duifje"
	phraseFrench               string = "h-h-h-hon"
	phraseGerman               string = "sssweetie"
	phraseItalian              string = "a-a-amor"
	phraseJapanese             string = "ウフフ"
	phraseLatinAmericanSpanish string = "ki-ki-kiú"
	phraseKorean               string = "우훗훗"
	phraseRussian              string = "м-милашка"
	phraseSpanish              string = "crustáceo"
	phraseSimplifiedChinese    string = "唔呼呼"
	phraseTraditionalChinese   string = "唔呼呼"
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
