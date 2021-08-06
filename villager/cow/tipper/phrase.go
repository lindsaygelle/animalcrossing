package tipper

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pushy"
	phraseCanadianFrench       string = "choupette"
	phraseDutch                string = "boelala"
	phraseFrench               string = "choupette"
	phraseGerman               string = "dingdong"
	phraseItalian              string = "sì sì sì"
	phraseJapanese             string = "ミルミル"
	phraseLatinAmericanSpanish string = "muamuuu"
	phraseKorean               string = "우유우유"
	phraseRussian              string = "теленок"
	phraseSpanish              string = "cuernitos"
	phraseSimplifiedChinese    string = "牛奶"
	phraseTraditionalChinese   string = "牛奶"
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
