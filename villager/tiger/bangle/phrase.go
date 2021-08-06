package bangle

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "growf"
	phraseCanadianFrench       string = "ma souris"
	phraseDutch                string = "grommf"
	phraseFrench               string = "ma souris"
	phraseGerman               string = "schnurf"
	phraseItalian              string = "grouf"
	phraseJapanese             string = "なのぉー"
	phraseLatinAmericanSpanish string = "gri-gruá"
	phraseKorean               string = "쪼옥쪽"
	phraseRussian              string = "р-р-рф"
	phraseSpanish              string = "mi vida"
	phraseSimplifiedChinese    string = "是喔"
	phraseTraditionalChinese   string = "是喔"
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
