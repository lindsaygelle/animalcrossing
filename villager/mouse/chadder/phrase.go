package chadder

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "fromage"
	phraseCanadianFrench       string = "cheese"
	phraseDutch                string = "cheese"
	phraseFrench               string = "frometon"
	phraseGerman               string = "hoho"
	phraseItalian              string = "iiik"
	phraseJapanese             string = "まあね"
	phraseLatinAmericanSpanish string = "fundifún"
	phraseKorean               string = "꼬리꼬리"
	phraseRussian              string = "сырок"
	phraseSpanish              string = "fundifún"
	phraseSimplifiedChinese    string = "也是啦"
	phraseTraditionalChinese   string = "也是啦"
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
