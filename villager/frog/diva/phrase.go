package diva

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ya know"
	phraseCanadianFrench       string = "cuicuisse"
	phraseDutch                string = "jeweetwel"
	phraseFrench               string = "cuicuisse"
	phraseGerman               string = "platsch"
	phraseItalian              string = "crabum"
	phraseJapanese             string = "ハーン"
	phraseLatinAmericanSpanish string = "crocró"
	phraseKorean               string = "흐응"
	phraseRussian              string = "знаешь"
	phraseSpanish              string = "crocró"
	phraseSimplifiedChinese    string = "蛤"
	phraseTraditionalChinese   string = "蛤"
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
