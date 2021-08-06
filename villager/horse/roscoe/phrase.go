package roscoe

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nay"
	phraseCanadianFrench       string = "shérif"
	phraseDutch                string = "nihihiet"
	phraseFrench               string = "shérif"
	phraseGerman               string = "galoppp"
	phraseItalian              string = "nay"
	phraseJapanese             string = "ブルル"
	phraseLatinAmericanSpanish string = "ñiii"
	phraseKorean               string = "부르르"
	phraseRussian              string = "не-е-ет"
	phraseSpanish              string = "bípedo"
	phraseSimplifiedChinese    string = "布鲁鲁"
	phraseTraditionalChinese   string = "布魯魯"
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
