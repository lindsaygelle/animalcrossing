package jitters

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bzzert"
	phraseCanadianFrench       string = "bouerk"
	phraseDutch                string = "tjulp"
	phraseFrench               string = "bouerk"
	phraseGerman               string = "zischhh"
	phraseItalian              string = "bzzert"
	phraseJapanese             string = "だニョ"
	phraseLatinAmericanSpanish string = "pío-pío"
	phraseKorean               string = "삼바"
	phraseRussian              string = "бзе-е-ерт"
	phraseSpanish              string = "pío-pío"
	phraseSimplifiedChinese    string = "诺诺"
	phraseTraditionalChinese   string = "諾諾"
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
