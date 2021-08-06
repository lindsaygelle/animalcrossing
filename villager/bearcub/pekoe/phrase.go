package pekoe

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bud"
	phraseCanadianFrench       string = "linou"
	phraseDutch                string = "bloempje"
	phraseFrench               string = "linou"
	phraseGerman               string = "flüster"
	phraseItalian              string = "purrrimba"
	phraseJapanese             string = "チャ"
	phraseLatinAmericanSpanish string = "pitiminí"
	phraseKorean               string = "띵호아"
	phraseRussian              string = "цветочек"
	phraseSpanish              string = "pitiminí"
	phraseSimplifiedChinese    string = "佳佳"
	phraseTraditionalChinese   string = "佳佳"
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
