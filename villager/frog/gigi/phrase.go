package gigi

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ribbette"
	phraseCanadianFrench       string = "crôazou"
	phraseDutch                string = "kwaakette"
	phraseFrench               string = "crôazou"
	phraseGerman               string = "quaki"
	phraseItalian              string = "gurghigù"
	phraseJapanese             string = "ベイビィ"
	phraseLatinAmericanSpanish string = "croá-croá"
	phraseKorean               string = "베이비"
	phraseRussian              string = "квакет"
	phraseSpanish              string = "trebolito"
	phraseSimplifiedChinese    string = "宝贝"
	phraseTraditionalChinese   string = "寶貝"
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
