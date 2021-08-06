package bud

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "shredded"
	phraseCanadianFrench       string = "mon man"
	phraseDutch                string = "maaaan"
	phraseFrench               string = "m'pote"
	phraseGerman               string = "kumpel"
	phraseItalian              string = "ruggi"
	phraseJapanese             string = "メ～ン"
	phraseLatinAmericanSpanish string = "grrrau"
	phraseKorean               string = "요～맨"
	phraseRussian              string = "че-е-ел"
	phraseSpanish              string = "sabes"
	phraseSimplifiedChinese    string = "男人"
	phraseTraditionalChinese   string = "男人"
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
