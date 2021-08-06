package cyd

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "rockin'"
	phraseCanadianFrench       string = "barda"
	phraseDutch                string = "mieters"
	phraseFrench               string = "bastringue"
	phraseGerman               string = "hau"
	phraseItalian              string = "rocking"
	phraseJapanese             string = "ウス"
	phraseLatinAmericanSpanish string = "rocanrol"
	phraseKorean               string = "피스"
	phraseRussian              string = "рок-н-ролл"
	phraseSpanish              string = "rocanrol"
	phraseSimplifiedChinese    string = "摇滚"
	phraseTraditionalChinese   string = "搖滾"
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
