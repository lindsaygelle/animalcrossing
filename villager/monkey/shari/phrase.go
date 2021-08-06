package shari

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cheeky"
	phraseCanadianFrench       string = "tavu"
	phraseDutch                string = "brutaaltje"
	phraseFrench               string = "tavu"
	phraseGerman               string = "makimaki"
	phraseItalian              string = "babbuh"
	phraseJapanese             string = "ウキキ"
	phraseLatinAmericanSpanish string = "babuí"
	phraseKorean               string = "우끼네"
	phraseRussian              string = "щекасто"
	phraseSpanish              string = "tunante"
	phraseSimplifiedChinese    string = "呜吱吱"
	phraseTraditionalChinese   string = "嗚吱吱"
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
