package mint

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ahhhhhh"
	phraseCanadianFrench       string = "ahhhhhh"
	phraseDutch                string = "ahhhhhh"
	phraseFrench               string = "ahhhhhh"
	phraseGerman               string = "knabber"
	phraseItalian              string = "squercia"
	phraseJapanese             string = "うっふん"
	phraseLatinAmericanSpanish string = "bellotuá"
	phraseKorean               string = "우훗"
	phraseRussian              string = "ах-ах"
	phraseSpanish              string = "bichito"
	phraseSimplifiedChinese    string = "喔哼"
	phraseTraditionalChinese   string = "喔哼"
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
