package monty

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "g'tang"
	phraseCanadianFrench       string = "l'outang"
	phraseDutch                string = "orang"
	phraseFrench               string = "l'outang"
	phraseGerman               string = "abbo"
	phraseItalian              string = "g'tang"
	phraseJapanese             string = "バナーナ"
	phraseLatinAmericanSpanish string = "uuuuh-ah"
	phraseKorean               string = "바나나"
	phraseRussian              string = "ух-эх"
	phraseSpanish              string = "uuuuh-ah"
	phraseSimplifiedChinese    string = "香蕉"
	phraseTraditionalChinese   string = "香蕉"
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
