package bluebear

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "peach"
	phraseCanadianFrench       string = "prune"
	phraseDutch                string = "perzik"
	phraseFrench               string = "prune"
	phraseGerman               string = "früchtchen"
	phraseItalian              string = "pepè"
	phraseJapanese             string = "キュン"
	phraseLatinAmericanSpanish string = "melí-melá"
	phraseKorean               string = "두근"
	phraseRussian              string = "персик"
	phraseSpanish              string = "cielito"
	phraseSimplifiedChinese    string = "怦怦"
	phraseTraditionalChinese   string = "怦怦"
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
