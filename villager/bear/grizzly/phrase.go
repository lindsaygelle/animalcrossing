package grizzly

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "grrr"
	phraseCanadianFrench       string = "canaillou"
	phraseDutch                string = "grrr"
	phraseFrench               string = "canaillou"
	phraseGerman               string = "grrr"
	phraseItalian              string = "grrr"
	phraseJapanese             string = "だクマ"
	phraseLatinAmericanSpanish string = "grrre"
	phraseKorean               string = "냅둬"
	phraseRussian              string = "грр"
	phraseSpanish              string = "tostón"
	phraseSimplifiedChinese    string = "熊熊"
	phraseTraditionalChinese   string = "熊熊"
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
