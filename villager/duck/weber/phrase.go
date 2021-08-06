package weber

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "quaa"
	phraseCanadianFrench       string = "plumeau"
	phraseDutch                string = "kwaaah"
	phraseFrench               string = "plumeau"
	phraseGerman               string = "quaa"
	phraseItalian              string = "quaaa"
	phraseJapanese             string = "ピヨ"
	phraseLatinAmericanSpanish string = "cuaa-cuaa"
	phraseKorean               string = "꽥꽥"
	phraseRussian              string = "кря-а"
	phraseSpanish              string = "cuaa-cuaa"
	phraseSimplifiedChinese    string = "唧唧"
	phraseTraditionalChinese   string = "唧唧"
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
