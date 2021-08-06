package claudia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ooh la la"
	phraseCanadianFrench       string = "grou grou"
	phraseDutch                string = "olala"
	phraseFrench               string = "grou grou"
	phraseGerman               string = "kchhh"
	phraseItalian              string = "smack"
	phraseJapanese             string = "アハ～ン"
	phraseLatinAmericanSpanish string = "mindundi"
	phraseKorean               string = "잇힝"
	phraseRussian              string = "о-ля-ля"
	phraseSpanish              string = "mindundi"
	phraseSimplifiedChinese    string = "啊哈"
	phraseTraditionalChinese   string = "啊哈"
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
