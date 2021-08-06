package diana

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "no doy"
	phraseCanadianFrench       string = "bichouchou"
	phraseDutch                string = "joh"
	phraseFrench               string = "bichouchou"
	phraseGerman               string = "kitz"
	phraseItalian              string = "mamy"
	phraseJapanese             string = "でしょ"
	phraseLatinAmericanSpanish string = "nonuá"
	phraseKorean               string = "라니까"
	phraseRussian              string = "еще бы"
	phraseSpanish              string = "braaavo"
	phraseSimplifiedChinese    string = "是吧"
	phraseTraditionalChinese   string = "是吧"
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
