package ike

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "roadie"
	phraseCanadianFrench       string = "pot d'miel"
	phraseDutch                string = "ja denk"
	phraseFrench               string = "pot d'miel"
	phraseGerman               string = "brrruah"
	phraseItalian              string = "sgrugno"
	phraseJapanese             string = "ボウズ"
	phraseLatinAmericanSpanish string = "grurrr"
	phraseKorean               string = "터프"
	phraseRussian              string = "гарр-сон"
	phraseSpanish              string = "tornillos"
	phraseSimplifiedChinese    string = "施主"
	phraseTraditionalChinese   string = "施主"
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
