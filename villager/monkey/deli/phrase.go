package deli

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "monch"
	phraseCanadianFrench       string = "liaaaane"
	phraseDutch                string = "smak"
	phraseFrench               string = "liaaaane"
	phraseGerman               string = "miekmiek"
	phraseItalian              string = "panzé"
	phraseJapanese             string = "だぜ"
	phraseLatinAmericanSpanish string = "ñim-ñim"
	phraseKorean               string = "흠냐리"
	phraseRussian              string = "чавк-чавк"
	phraseSpanish              string = "millón"
	phraseSimplifiedChinese    string = "去吧"
	phraseTraditionalChinese   string = "去吧"
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
