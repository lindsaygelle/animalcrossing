package charlise

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "urgh"
	phraseCanadianFrench       string = "chacha"
	phraseDutch                string = "pff"
	phraseFrench               string = "verlue"
	phraseGerman               string = "bärbär"
	phraseItalian              string = "sbuff"
	phraseJapanese             string = "よいしょ"
	phraseLatinAmericanSpanish string = "urrrgh"
	phraseKorean               string = "으랏차차"
	phraseRussian              string = "ого-го"
	phraseSpanish              string = "comolooyes"
	phraseSimplifiedChinese    string = "唷咻"
	phraseTraditionalChinese   string = "唷咻"
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
