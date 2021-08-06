package freya

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "uff da"
	phraseCanadianFrench       string = "mon loulou"
	phraseDutch                string = "fjord"
	phraseFrench               string = "mon loulou"
	phraseGerman               string = "ähmmm"
	phraseItalian              string = "uuuff"
	phraseJapanese             string = "なのだわ"
	phraseLatinAmericanSpanish string = "auuuff"
	phraseKorean               string = "유행이야"
	phraseRussian              string = "хрусть"
	phraseSpanish              string = "auuuff"
	phraseSimplifiedChinese    string = "就是说哇"
	phraseTraditionalChinese   string = "就是說哇"
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
