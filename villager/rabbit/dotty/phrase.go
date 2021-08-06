package dotty

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "wee one"
	phraseCanadianFrench       string = "lapinou"
	phraseDutch                string = "lamprei"
	phraseFrench               string = "lapinou"
	phraseGerman               string = "uiui"
	phraseItalian              string = "caroté"
	phraseJapanese             string = "ラン"
	phraseLatinAmericanSpanish string = "toini"
	phraseKorean               string = "룰루랄라"
	phraseRussian              string = "малыш"
	phraseSpanish              string = "trompis"
	phraseSimplifiedChinese    string = "啷"
	phraseTraditionalChinese   string = "啷"
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
