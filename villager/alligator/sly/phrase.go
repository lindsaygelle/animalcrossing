package sly

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hoo-rah"
	phraseCanadianFrench       string = "repos"
	phraseDutch                string = "ingerukt"
	phraseFrench               string = "repos"
	phraseGerman               string = "schnapp"
	phraseItalian              string = "ciaf"
	phraseJapanese             string = "カサコソ"
	phraseLatinAmericanSpanish string = "smack"
	phraseKorean               string = "부스럭"
	phraseRussian              string = "ура-а-а"
	phraseSpanish              string = "capón"
	phraseSimplifiedChinese    string = "唰唰"
	phraseTraditionalChinese   string = "唰唰"
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
