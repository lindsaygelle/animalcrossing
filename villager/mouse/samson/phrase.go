package samson

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pipsqueak"
	phraseCanadianFrench       string = "scouic"
	phraseDutch                string = "piepklein"
	phraseFrench               string = "scouic"
	phraseGerman               string = "pieeeps"
	phraseItalian              string = "pisquak"
	phraseJapanese             string = "チュー"
	phraseLatinAmericanSpanish string = "escuác"
	phraseKorean               string = "쪽"
	phraseRussian              string = "ревушка"
	phraseSpanish              string = "flojeras"
	phraseSimplifiedChinese    string = "吱"
	phraseTraditionalChinese   string = "吱"
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
