package bam

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "kablang"
	phraseCanadianFrench       string = "têtedebois"
	phraseDutch                string = "gaffel"
	phraseFrench               string = "têtedebois"
	phraseGerman               string = "weihweih"
	phraseItalian              string = "hop hop"
	phraseJapanese             string = "ナラね"
	phraseLatinAmericanSpanish string = "caplúúú"
	phraseKorean               string = "슉슉"
	phraseRussian              string = "скок-скок"
	phraseSpanish              string = "galilla"
	phraseSimplifiedChinese    string = "梅花"
	phraseTraditionalChinese   string = "梅花"
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
