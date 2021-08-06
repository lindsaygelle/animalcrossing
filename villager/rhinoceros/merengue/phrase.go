package merengue

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "shortcake"
	phraseCanadianFrench       string = "confiture"
	phraseDutch                string = "schuimpje"
	phraseFrench               string = "ma fraise"
	phraseGerman               string = "trarapp"
	phraseItalian              string = "rien"
	phraseJapanese             string = "でシュガ"
	phraseLatinAmericanSpanish string = "fruti"
	phraseKorean               string = "베리베리"
	phraseRussian              string = "кексик"
	phraseSpanish              string = "fresitas"
	phraseSimplifiedChinese    string = "糖糖"
	phraseTraditionalChinese   string = "糖糖"
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
