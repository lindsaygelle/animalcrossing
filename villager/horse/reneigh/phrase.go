package reneigh

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ayup yup"
	phraseCanadianFrench       string = "crincrin"
	phraseDutch                string = "uhuh"
	phraseFrench               string = "crincrin"
	phraseGerman               string = "tada"
	phraseItalian              string = "hiii-op"
	phraseJapanese             string = "スン"
	phraseLatinAmericanSpanish string = "juááá"
	phraseKorean               string = "쉬익"
	phraseRussian              string = "но-но"
	phraseSpanish              string = "alehop"
	phraseSimplifiedChinese    string = "哼"
	phraseTraditionalChinese   string = "哼"
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
