package blaire

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nutlet"
	phraseCanadianFrench       string = "coquille"
	phraseDutch                string = "notendop"
	phraseFrench               string = "coquille"
	phraseGerman               string = "knack"
	phraseItalian              string = "scrick"
	phraseJapanese             string = "ふふん"
	phraseLatinAmericanSpanish string = "ñicuá"
	phraseKorean               string = "헤헤헤"
	phraseRussian              string = "крошка"
	phraseSpanish              string = "bayita"
	phraseSimplifiedChinese    string = "哼哼"
	phraseTraditionalChinese   string = "哼哼"
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
