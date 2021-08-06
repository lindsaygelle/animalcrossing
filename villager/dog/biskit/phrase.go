package biskit

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "dawg"
	phraseCanadianFrench       string = "clair"
	phraseDutch                string = "makker"
	phraseFrench               string = "clair"
	phraseGerman               string = "kläff"
	phraseItalian              string = "brrranco"
	phraseJapanese             string = "だイヌ"
	phraseLatinAmericanSpanish string = "guau"
	phraseKorean               string = "멍멍"
	phraseRussian              string = "пр-риятяв"
	phraseSpanish              string = "guau"
	phraseSimplifiedChinese    string = "狗狗"
	phraseTraditionalChinese   string = "狗狗"
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
