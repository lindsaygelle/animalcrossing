package goldie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "woof"
	phraseCanadianFrench       string = "ouaf"
	phraseDutch                string = "woef"
	phraseFrench               string = "ouaf"
	phraseGerman               string = "wuff"
	phraseItalian              string = "wuf wuf"
	phraseJapanese             string = "ワン"
	phraseLatinAmericanSpanish string = "guau-guau"
	phraseKorean               string = "왈왈"
	phraseRussian              string = "тяв"
	phraseSpanish              string = "guau-guau"
	phraseSimplifiedChinese    string = "汪"
	phraseTraditionalChinese   string = "汪"
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
