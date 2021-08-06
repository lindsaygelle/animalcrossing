package tammy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ya heard"
	phraseCanadianFrench       string = "grifouille"
	phraseDutch                string = "weet je"
	phraseFrench               string = "grifouille"
	phraseGerman               string = "bossi"
	phraseItalian              string = "ohè"
	phraseJapanese             string = "だってサ"
	phraseLatinAmericanSpanish string = "tararí"
	phraseKorean               string = "차라리"
	phraseRussian              string = "понимаешь"
	phraseSpanish              string = "tararí"
	phraseSimplifiedChinese    string = "竟然说"
	phraseTraditionalChinese   string = "竟然說"
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
