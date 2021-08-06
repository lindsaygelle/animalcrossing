package walker

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "wuh"
	phraseCanadianFrench       string = "ouh"
	phraseDutch                string = "waf"
	phraseFrench               string = "ouh"
	phraseGerman               string = "wafff"
	phraseItalian              string = "wuff"
	phraseJapanese             string = "バウ"
	phraseLatinAmericanSpanish string = "guuu-ah"
	phraseKorean               string = "컹컹"
	phraseRussian              string = "ау-у"
	phraseSpanish              string = "huesitos"
	phraseSimplifiedChinese    string = "咆"
	phraseTraditionalChinese   string = "咆"
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
