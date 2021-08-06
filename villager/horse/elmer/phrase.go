package elmer

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tenderfoot"
	phraseCanadianFrench       string = "mouuuuiii"
	phraseDutch                string = "jonkie"
	phraseFrench               string = "mouuuuiii"
	phraseGerman               string = "anfänger"
	phraseItalian              string = "pinopino"
	phraseJapanese             string = "だヒヒン"
	phraseLatinAmericanSpanish string = "hop-uah"
	phraseKorean               string = "히히힝"
	phraseRussian              string = "жеребенок"
	phraseSpanish              string = "pezuñitas"
	phraseSimplifiedChinese    string = "酥酥"
	phraseTraditionalChinese   string = "酥酥"
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
