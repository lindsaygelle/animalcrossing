package mathilda

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "wee baby"
	phraseCanadianFrench       string = "riquiqui"
	phraseDutch                string = "ukkepuk"
	phraseFrench               string = "riquiqui"
	phraseGerman               string = "knirps"
	phraseItalian              string = "oplà"
	phraseJapanese             string = "ッハ"
	phraseLatinAmericanSpanish string = "ala-rorro"
	phraseKorean               string = "응"
	phraseRussian              string = "малютка"
	phraseSpanish              string = "peso mosca"
	phraseSimplifiedChinese    string = "嘿"
	phraseTraditionalChinese   string = "嘿"
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
