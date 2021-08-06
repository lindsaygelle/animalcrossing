package biff

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "squirt"
	phraseCanadianFrench       string = "hips hips"
	phraseDutch                string = "spetter"
	phraseFrench               string = "hips hips"
	phraseGerman               string = "skwiirt"
	phraseItalian              string = "squirt"
	phraseJapanese             string = "じゃけん"
	phraseLatinAmericanSpanish string = "hipo-pipo"
	phraseKorean               string = "오도독"
	phraseRussian              string = "малек"
	phraseSpanish              string = "hipo-pipo"
	phraseSimplifiedChinese    string = "因为"
	phraseTraditionalChinese   string = "因為"
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
