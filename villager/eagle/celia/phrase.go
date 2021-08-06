package celia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "feathers"
	phraseCanadianFrench       string = "bebec"
	phraseDutch                string = "pluimpje"
	phraseFrench               string = "bebec"
	phraseGerman               string = "kraaack"
	phraseItalian              string = "ah bon"
	phraseJapanese             string = "そうね"
	phraseLatinAmericanSpanish string = "plumí"
	phraseKorean               string = "그래맞아"
	phraseRussian              string = "перышко"
	phraseSpanish              string = "plumis"
	phraseSimplifiedChinese    string = "是呢"
	phraseTraditionalChinese   string = "是呢"
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
