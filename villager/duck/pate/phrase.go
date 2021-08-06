package pate

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "quackle"
	phraseCanadianFrench       string = "coin coin"
	phraseDutch                string = "kwakel"
	phraseFrench               string = "coin coin"
	phraseGerman               string = "quaak"
	phraseItalian              string = "quackquack"
	phraseJapanese             string = "メソメソ"
	phraseLatinAmericanSpanish string = "cuaqui"
	phraseKorean               string = "아잉"
	phraseRussian              string = "кряки-кряк"
	phraseSpanish              string = "cuaqui"
	phraseSimplifiedChinese    string = "呜呜"
	phraseTraditionalChinese   string = "嗚嗚"
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
