package megan

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sundae"
	phraseCanadianFrench       string = "berlingot"
	phraseDutch                string = "lolly"
	phraseFrench               string = "berlingot"
	phraseGerman               string = "lollipop"
	phraseItalian              string = "sucré"
	phraseJapanese             string = "ぺろ"
	phraseLatinAmericanSpanish string = "tutifruti"
	phraseKorean               string = "낼름"
	phraseRussian              string = "сладенько"
	phraseSpanish              string = "tutifruti"
	phraseSimplifiedChinese    string = "舔舔"
	phraseTraditionalChinese   string = "舔舔"
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
