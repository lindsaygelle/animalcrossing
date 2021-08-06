package inkwell

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "splat"
	phraseCanadianFrench       string = "splat"
	phraseDutch                string = ""
	phraseFrench               string = "splat"
	phraseGerman               string = "splat"
	phraseItalian              string = "splat"
	phraseJapanese             string = "ガチで"
	phraseLatinAmericanSpanish string = "tinta va"
	phraseKorean               string = "진심"
	phraseRussian              string = ""
	phraseSpanish              string = "tinta va"
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
