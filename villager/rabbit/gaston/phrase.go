package gaston

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mon chou"
	phraseCanadianFrench       string = "cahuète"
	phraseDutch                string = "mon amour"
	phraseFrench               string = "cahuète"
	phraseGerman               string = "herzblatt"
	phraseItalian              string = "cherie"
	phraseJapanese             string = "ムホッ"
	phraseLatinAmericanSpanish string = "mon chou"
	phraseKorean               string = "쿨럭"
	phraseRussian              string = "мон зайшу"
	phraseSpanish              string = "mon chou"
	phraseSimplifiedChinese    string = "唔贺"
	phraseTraditionalChinese   string = "唔賀"
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
