package filly

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hah"
	phraseCanadianFrench       string = "papaye"
	phraseDutch                string = ""
	phraseFrench               string = "papaye"
	phraseGerman               string = "hühott"
	phraseItalian              string = "galop"
	phraseJapanese             string = "はぁっ！"
	phraseLatinAmericanSpanish string = "galope"
	phraseKorean               string = "타앗"
	phraseRussian              string = ""
	phraseSpanish              string = "galope"
	phraseSimplifiedChinese    string = ""
	phraseTraditionalChinese   string = ""
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
