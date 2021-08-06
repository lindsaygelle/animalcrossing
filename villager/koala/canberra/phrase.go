package canberra

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nuh uh"
	phraseCanadianFrench       string = "zyva"
	phraseDutch                string = "echnie"
	phraseFrench               string = "zyva"
	phraseGerman               string = "kolala"
	phraseItalian              string = "eurekalì"
	phraseJapanese             string = "え～"
	phraseLatinAmericanSpanish string = "koalalá"
	phraseKorean               string = "어어"
	phraseRussian              string = "ох-хо"
	phraseSpanish              string = "kolalá"
	phraseSimplifiedChinese    string = "咦"
	phraseTraditionalChinese   string = "咦～"
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
