package sydney

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sunshine"
	phraseCanadianFrench       string = "mon trésor"
	phraseDutch                string = "zonnetje"
	phraseFrench               string = "mon trésor"
	phraseGerman               string = "sternchen"
	phraseItalian              string = "calipto"
	phraseJapanese             string = "だコアラ"
	phraseLatinAmericanSpanish string = "calipto"
	phraseKorean               string = "발그레"
	phraseRussian              string = "солнышко"
	phraseSpanish              string = "súper"
	phraseSimplifiedChinese    string = "无尾熊"
	phraseTraditionalChinese   string = "無尾熊"
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
