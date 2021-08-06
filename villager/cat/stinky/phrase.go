package stinky

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "GAAHHH"
	phraseCanadianFrench       string = "miaoum"
	phraseDutch                string = "BAH"
	phraseFrench               string = "miaoum"
	phraseGerman               string = "fauchch"
	phraseItalian              string = "GAAHHH"
	phraseJapanese             string = "ダァーッ"
	phraseLatinAmericanSpanish string = "marramiau"
	phraseKorean               string = "아뵤"
	phraseRussian              string = "фу-у-у"
	phraseSpanish              string = "marramiau"
	phraseSimplifiedChinese    string = "打啊"
	phraseTraditionalChinese   string = "打啊"
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
