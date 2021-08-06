package jambette

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "croak-kay"
	phraseCanadianFrench       string = "quoi"
	phraseDutch                string = "inderkwaak"
	phraseFrench               string = "quoi"
	phraseGerman               string = "quaacki"
	phraseItalian              string = "crok-ok"
	phraseJapanese             string = "ですわよ"
	phraseLatinAmericanSpanish string = "croac"
	phraseKorean               string = "쪼오옥"
	phraseRussian              string = "кво-кей"
	phraseSpanish              string = "croac"
	phraseSimplifiedChinese    string = "是蛙"
	phraseTraditionalChinese   string = "是蛙"
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
