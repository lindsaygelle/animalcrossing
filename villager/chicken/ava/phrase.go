package ava

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "beaker"
	phraseCanadianFrench       string = "cô-cômment"
	phraseDutch                string = "snavel"
	phraseFrench               string = "cô-cômment"
	phraseGerman               string = "biiika"
	phraseItalian              string = "sbecsbec"
	phraseJapanese             string = "のよぉ"
	phraseLatinAmericanSpanish string = "kikiú"
	phraseKorean               string = "그랬대요"
	phraseRussian              string = "клювик"
	phraseSpanish              string = "piquito"
	phraseSimplifiedChinese    string = "是唷"
	phraseTraditionalChinese   string = "是唷"
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
