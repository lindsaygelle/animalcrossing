package moe

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "myawn"
	phraseCanadianFrench       string = "miou"
	phraseDutch                string = "geeuweldig"
	phraseFrench               string = "miou"
	phraseGerman               string = "schluchz"
	phraseItalian              string = "myawn"
	phraseJapanese             string = "な"
	phraseLatinAmericanSpanish string = "grrrroou"
	phraseKorean               string = "그치"
	phraseRussian              string = "мяу-а-а-а"
	phraseSpanish              string = "grrrroou"
	phraseSimplifiedChinese    string = "呐"
	phraseTraditionalChinese   string = "吶"
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
