package frita

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "oh ewe"
	phraseCanadianFrench       string = "mouêêê"
	phraseDutch                string = "ooi zeg"
	phraseFrench               string = "mouêêê"
	phraseGerman               string = "knuddel"
	phraseItalian              string = "beee la la"
	phraseJapanese             string = "だポテ"
	phraseLatinAmericanSpanish string = "berebé"
	phraseKorean               string = "포테토"
	phraseRussian              string = "мутончик"
	phraseSpanish              string = "berebé"
	phraseSimplifiedChinese    string = "薯薯"
	phraseTraditionalChinese   string = "薯薯"
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
