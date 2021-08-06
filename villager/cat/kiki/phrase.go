package kiki

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "kitty cat"
	phraseCanadianFrench       string = "chti minou"
	phraseDutch                string = "poesjemauw"
	phraseFrench               string = "chti minou"
	phraseGerman               string = "miausi"
	phraseItalian              string = "mieaoo"
	phraseJapanese             string = "だニ"
	phraseLatinAmericanSpanish string = "michifiú"
	phraseKorean               string = "뭐라지요"
	phraseRussian              string = "кисуля"
	phraseSpanish              string = "galletita"
	phraseSimplifiedChinese    string = "酱酱"
	phraseTraditionalChinese   string = "醬醬"
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
