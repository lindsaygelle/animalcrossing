package iggly

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "waddler"
	phraseCanadianFrench       string = "zozio"
	phraseDutch                string = "waggelaar"
	phraseFrench               string = "zozio"
	phraseGerman               string = "plitsch"
	phraseItalian              string = "tubi tubi"
	phraseJapanese             string = "クルクル"
	phraseLatinAmericanSpanish string = "chirpi"
	phraseKorean               string = "펭글펭글"
	phraseRussian              string = "ластошлеп"
	phraseSpanish              string = "chirpi"
	phraseSimplifiedChinese    string = "卷卷"
	phraseTraditionalChinese   string = "捲捲"
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
