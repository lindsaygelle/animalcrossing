package coco

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "doyoing"
	phraseCanadianFrench       string = "touloulou"
	phraseDutch                string = "bojoing"
	phraseFrench               string = "touloulou"
	phraseGerman               string = "hoppsa"
	phraseItalian              string = "evviva"
	phraseJapanese             string = "はにょ"
	phraseLatinAmericanSpanish string = "toin-toin"
	phraseKorean               string = "삐용"
	phraseRussian              string = "прыг-скок"
	phraseSpanish              string = "toin-toin"
	phraseSimplifiedChinese    string = "土俑"
	phraseTraditionalChinese   string = "土俑"
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
