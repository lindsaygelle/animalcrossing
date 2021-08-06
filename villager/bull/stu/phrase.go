package stu

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mrooooo"
	phraseCanadianFrench       string = "meuh quoi"
	phraseDutch                string = "boe-jend"
	phraseFrench               string = "meuh quoi"
	phraseGerman               string = "muhuhu"
	phraseItalian              string = "ciaumù"
	phraseJapanese             string = "だなっす"
	phraseLatinAmericanSpanish string = "olé y olé"
	phraseKorean               string = "알것소"
	phraseRussian              string = "му-у"
	phraseSpanish              string = "olé y olé"
	phraseSimplifiedChinese    string = "奔驰"
	phraseTraditionalChinese   string = "奔馳"
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
