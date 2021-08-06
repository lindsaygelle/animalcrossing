package felyne

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nya"
	phraseCanadianFrench       string = "miaou"
	phraseDutch                string = ""
	phraseFrench               string = "miaou"
	phraseGerman               string = "miau"
	phraseItalian              string = "mià"
	phraseJapanese             string = "ニャ"
	phraseLatinAmericanSpanish string = "miaumo"
	phraseKorean               string = "냥"
	phraseRussian              string = ""
	phraseSpanish              string = "miaumo"
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
