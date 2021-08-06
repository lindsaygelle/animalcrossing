package angus

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "macmoo"
	phraseCanadianFrench       string = "meuh meuh"
	phraseDutch                string = "denderund"
	phraseFrench               string = "meuh meuh"
	phraseGerman               string = "brrruzzl"
	phraseItalian              string = "macmoo"
	phraseJapanese             string = "ふふ"
	phraseLatinAmericanSpanish string = "bufff"
	phraseKorean               string = "후후"
	phraseRussian              string = "макму-у-у"
	phraseSpanish              string = "cuernos"
	phraseSimplifiedChinese    string = "呼呼"
	phraseTraditionalChinese   string = "呼呼"
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
