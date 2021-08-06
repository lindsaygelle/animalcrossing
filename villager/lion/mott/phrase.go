package mott

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cagey"
	phraseCanadianFrench       string = "grif-grif"
	phraseDutch                string = "rrrrrauw"
	phraseFrench               string = "grif-grif"
	phraseGerman               string = "gröhl"
	phraseItalian              string = "groarrivo"
	phraseJapanese             string = "ハハハ"
	phraseLatinAmericanSpanish string = "sapristi"
	phraseKorean               string = "하하하"
	phraseRussian              string = "зоосад"
	phraseSpanish              string = "sapristi"
	phraseSimplifiedChinese    string = "哈哈哈"
	phraseTraditionalChinese   string = "哈哈哈"
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
