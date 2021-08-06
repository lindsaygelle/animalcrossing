package punchy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mrmpht"
	phraseCanadianFrench       string = "mimine"
	phraseDutch                string = "plof"
	phraseFrench               string = "blébléblé"
	phraseGerman               string = "mrmpft"
	phraseItalian              string = "maramao"
	phraseJapanese             string = "だのら"
	phraseLatinAmericanSpanish string = "gruuuah"
	phraseKorean               string = "노라줘"
	phraseRussian              string = "отдыхаем"
	phraseSpanish              string = "arenque"
	phraseSimplifiedChinese    string = "晃晃"
	phraseTraditionalChinese   string = "晃晃"
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
