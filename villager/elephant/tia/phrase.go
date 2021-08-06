package tia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "teacup"
	phraseCanadianFrench       string = "énoooorme"
	phraseDutch                string = "kopje thee"
	phraseFrench               string = "énoooorme"
	phraseGerman               string = "huiuiui"
	phraseItalian              string = "parbleu"
	phraseJapanese             string = "ホッ"
	phraseLatinAmericanSpanish string = "snuuuf"
	phraseKorean               string = "따끈따끈"
	phraseRussian              string = "чайку-у"
	phraseSpanish              string = "snuuuf"
	phraseSimplifiedChinese    string = "呼"
	phraseTraditionalChinese   string = "呼"
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
