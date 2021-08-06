package rory

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "capital"
	phraseCanadianFrench       string = "excellent"
	phraseDutch                string = "koninklijk"
	phraseFrench               string = "criniou"
	phraseGerman               string = "graaah"
	phraseItalian              string = "roar"
	phraseJapanese             string = "ナハッ"
	phraseLatinAmericanSpanish string = "ticotaco"
	phraseKorean               string = "크릉"
	phraseRussian              string = "капитально"
	phraseSpanish              string = "ticotaco"
	phraseSimplifiedChinese    string = "哪哈"
	phraseTraditionalChinese   string = "哪哈"
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
