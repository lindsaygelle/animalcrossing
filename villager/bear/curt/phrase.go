package curt

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "fuzzball"
	phraseCanadianFrench       string = "bouboule"
	phraseDutch                string = "brombeer"
	phraseFrench               string = "bouboule"
	phraseGerman               string = "grumml"
	phraseItalian              string = "grrroan"
	phraseJapanese             string = "ウム"
	phraseLatinAmericanSpanish string = "gromp"
	phraseKorean               string = "음"
	phraseRussian              string = "махрово"
	phraseSpanish              string = "oye"
	phraseSimplifiedChinese    string = "嗯唔"
	phraseTraditionalChinese   string = "嗯唔"
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
