package truffles

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snoutie"
	phraseCanadianFrench       string = "jambon"
	phraseDutch                string = "snuiter"
	phraseFrench               string = "jambon"
	phraseGerman               string = "schnauzie"
	phraseItalian              string = "ficcanaso"
	phraseJapanese             string = "だわさ"
	phraseLatinAmericanSpanish string = "oink-oink"
	phraseKorean               string = "거라예"
	phraseRussian              string = "хрюша"
	phraseSpanish              string = "oink-oink"
	phraseSimplifiedChinese    string = "哇赛"
	phraseTraditionalChinese   string = "哇賽"
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
