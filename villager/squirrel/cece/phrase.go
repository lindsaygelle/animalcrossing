package cece

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "stay fresh"
	phraseCanadianFrench       string = "maréôte"
	phraseDutch                string = ""
	phraseFrench               string = "maréôte"
	phraseGerman               string = "aioli"
	phraseItalian              string = "calacala"
	phraseJapanese             string = "よろしく"
	phraseLatinAmericanSpanish string = "cosmar"
	phraseKorean               string = "잘지내자"
	phraseRussian              string = ""
	phraseSpanish              string = "cosmar"
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
