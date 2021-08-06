package stitches

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "stuffin'"
	phraseCanadianFrench       string = "ou bien"
	phraseDutch                string = "pluche"
	phraseFrench               string = "ou bien"
	phraseGerman               string = "brummm"
	phraseItalian              string = "ohilà"
	phraseJapanese             string = "なのれす"
	phraseLatinAmericanSpanish string = "paguahhh"
	phraseKorean               string = "그런거죠"
	phraseRussian              string = "плюш-плюш"
	phraseSpanish              string = "peluche"
	phraseSimplifiedChinese    string = "玩玩"
	phraseTraditionalChinese   string = "玩玩"
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
