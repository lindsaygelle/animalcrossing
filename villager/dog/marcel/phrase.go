package marcel

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "non"
	phraseCanadianFrench       string = "miaouf"
	phraseDutch                string = "mais non"
	phraseFrench               string = "miaouf"
	phraseGerman               string = "tusch"
	phraseItalian              string = "snif snif"
	phraseJapanese             string = "なんじゃ"
	phraseLatinAmericanSpanish string = "nenonó"
	phraseKorean               string = "하옵소서"
	phraseRussian              string = "ау-уи"
	phraseSpanish              string = "colmillos"
	phraseSimplifiedChinese    string = "什么字"
	phraseTraditionalChinese   string = "什麼字"
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
