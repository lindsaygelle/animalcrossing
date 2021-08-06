package cleo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sugar"
	phraseCanadianFrench       string = "beauté"
	phraseDutch                string = "snoepje"
	phraseFrench               string = "beauté"
	phraseGerman               string = "plätzchen"
	phraseItalian              string = "dolcetto"
	phraseJapanese             string = "あそばせ"
	phraseLatinAmericanSpanish string = "hop-puá"
	phraseKorean               string = "훗훗"
	phraseRussian              string = "рафинад"
	phraseSpanish              string = "bomboncito"
	phraseSimplifiedChinese    string = "游游"
	phraseTraditionalChinese   string = "遊遊"
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
