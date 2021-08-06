package renee

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = ""
	phraseCanadianFrench       string = ""
	phraseDutch                string = ""
	phraseFrench               string = ""
	phraseGerman               string = ""
	phraseItalian              string = ""
	phraseJapanese             string = ""
	phraseLatinAmericanSpanish string = ""
	phraseKorean               string = ""
	phraseRussian              string = ""
	phraseSpanish              string = ""
	phraseSimplifiedChinese    string = ""
	phraseTraditionalChinese   string = ""
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
