package bea

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bingo"
	phraseCanadianFrench       string = "mon chaton"
	phraseDutch                string = "bingo"
	phraseFrench               string = "mon chaton"
	phraseGerman               string = "bingo"
	phraseItalian              string = "bingo"
	phraseJapanese             string = "グー"
	phraseLatinAmericanSpanish string = "gufi-guf"
	phraseKorean               string = "쫀득"
	phraseRussian              string = "динго"
	phraseSpanish              string = "huesín"
	phraseSimplifiedChinese    string = "宾果"
	phraseTraditionalChinese   string = "賓果"
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
