package spork

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snork"
	phraseCanadianFrench       string = "diantre"
	phraseDutch                string = "snurk"
	phraseFrench               string = "diantre"
	phraseGerman               string = "grunz"
	phraseItalian              string = "snork"
	phraseJapanese             string = "だブー"
	phraseLatinAmericanSpanish string = "pork"
	phraseKorean               string = "꼬르륵"
	phraseRussian              string = "хрю"
	phraseSpanish              string = "pork"
	phraseSimplifiedChinese    string = "噗"
	phraseTraditionalChinese   string = "噗"
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
