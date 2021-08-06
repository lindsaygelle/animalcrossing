package vesta

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "baaaffo"
	phraseCanadianFrench       string = "foulââârd"
	phraseDutch                string = "bèèè"
	phraseFrench               string = "pupull"
	phraseGerman               string = "bäääh"
	phraseItalian              string = "beeeby"
	phraseJapanese             string = "なのね"
	phraseLatinAmericanSpanish string = "beeeibi"
	phraseKorean               string = "쩝"
	phraseRussian              string = "бе-е"
	phraseSpanish              string = "beeeibi"
	phraseSimplifiedChinese    string = "就这样呀"
	phraseTraditionalChinese   string = "就這樣呀"
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
