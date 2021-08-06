package timbra

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pine nut"
	phraseCanadianFrench       string = "barbiche"
	phraseDutch                string = "denappel"
	phraseFrench               string = "barbiche"
	phraseGerman               string = "bwäääh"
	phraseItalian              string = "pecoramè"
	phraseJapanese             string = "まつわ"
	phraseLatinAmericanSpanish string = "yalovééés"
	phraseKorean               string = "달려"
	phraseRussian              string = "сосенка"
	phraseSpanish              string = "yalovéees"
	phraseSimplifiedChinese    string = "等啊"
	phraseTraditionalChinese   string = "等啊"
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
