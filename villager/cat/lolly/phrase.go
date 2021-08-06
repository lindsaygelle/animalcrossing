package lolly

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bonbon"
	phraseCanadianFrench       string = "ron-ron"
	phraseDutch                string = "zuurtje"
	phraseFrench               string = "ron-ron"
	phraseGerman               string = "kratz"
	phraseItalian              string = "purrrr"
	phraseJapanese             string = "あのね"
	phraseLatinAmericanSpanish string = "purrrr"
	phraseKorean               string = "퐁퐁"
	phraseRussian              string = "конфетка"
	phraseSpanish              string = "purrrr"
	phraseSimplifiedChinese    string = "娜个"
	phraseTraditionalChinese   string = "娜個"
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
