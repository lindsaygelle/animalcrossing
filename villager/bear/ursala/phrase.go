package ursala

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "grooomph"
	phraseCanadianFrench       string = "groumpf"
	phraseDutch                string = "oemf"
	phraseFrench               string = "groumpf"
	phraseGerman               string = "brumbrum"
	phraseItalian              string = "gruuunf"
	phraseJapanese             string = "やーねぇ"
	phraseLatinAmericanSpanish string = "grumf"
	phraseKorean               string = "그라지"
	phraseRussian              string = "гррум"
	phraseSpanish              string = "grumf"
	phraseSimplifiedChinese    string = "呀呐"
	phraseTraditionalChinese   string = "呀吶"
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
