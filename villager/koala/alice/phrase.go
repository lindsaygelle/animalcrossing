package alice

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "guvnor"
	phraseCanadianFrench       string = "chef"
	phraseDutch                string = "chef"
	phraseFrench               string = "chef"
	phraseGerman               string = "schnarchhh"
	phraseItalian              string = "tesorino"
	phraseJapanese             string = "キラリ"
	phraseLatinAmericanSpanish string = "eucaliup"
	phraseKorean               string = "반짝"
	phraseRussian              string = "лидер"
	phraseSpanish              string = "tú"
	phraseSimplifiedChinese    string = "晶亮"
	phraseTraditionalChinese   string = "晶亮"
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
