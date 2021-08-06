package henry

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snoozit"
	phraseCanadianFrench       string = "kôwa"
	phraseDutch                string = "snurkwaak"
	phraseFrench               string = "pustule"
	phraseGerman               string = "quappi"
	phraseItalian              string = "croac"
	phraseJapanese             string = "むにゃ"
	phraseLatinAmericanSpanish string = "cruacruá"
	phraseKorean               string = "음냐음냐"
	phraseRussian              string = "ква-храп"
	phraseSpanish              string = "cruacruá"
	phraseSimplifiedChinese    string = "滴咕"
	phraseTraditionalChinese   string = "滴咕"
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
