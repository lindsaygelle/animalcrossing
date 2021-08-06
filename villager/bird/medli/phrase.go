package medli

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "gimme"
	phraseCanadianFrench       string = "lalalyre"
	phraseDutch                string = ""
	phraseFrench               string = "lalalyre"
	phraseGerman               string = "hilfhilf"
	phraseItalian              string = "plinplon"
	phraseJapanese             string = "くらさい"
	phraseLatinAmericanSpanish string = "plimplín"
	phraseKorean               string = "핑그르르"
	phraseRussian              string = ""
	phraseSpanish              string = "plimplín"
	phraseSimplifiedChinese    string = ""
	phraseTraditionalChinese   string = ""
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
