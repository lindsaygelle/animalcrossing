package cobb

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hot dog"
	phraseCanadianFrench       string = "andouille"
	phraseDutch                string = "knakker"
	phraseFrench               string = "andouille"
	phraseGerman               string = "würstel"
	phraseItalian              string = "salsicce"
	phraseJapanese             string = "でアール"
	phraseLatinAmericanSpanish string = "saperlipop"
	phraseKorean               string = "이노라"
	phraseRussian              string = "сосиска"
	phraseSpanish              string = "saperlipop"
	phraseSimplifiedChinese    string = "公亩"
	phraseTraditionalChinese   string = "公畝"
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
