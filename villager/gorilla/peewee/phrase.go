package peewee

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "li'l bitty baby"
	phraseCanadianFrench       string = "gadget"
	phraseDutch                string = "dwerg"
	phraseFrench               string = "gadget"
	phraseGerman               string = "zwerg"
	phraseItalian              string = "baaanane"
	phraseJapanese             string = "ガオ"
	phraseLatinAmericanSpanish string = "cani"
	phraseKorean               string = "겉멋"
	phraseRussian              string = "мелкота"
	phraseSpanish              string = "cani"
	phraseSimplifiedChinese    string = "高高"
	phraseTraditionalChinese   string = "高高"
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
