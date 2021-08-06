package gala

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snortie"
	phraseCanadianFrench       string = "tigroin"
	phraseDutch                string = "snufferd"
	phraseFrench               string = "tigroin"
	phraseGerman               string = "lächel"
	phraseItalian              string = "gruffeffè"
	phraseJapanese             string = "ちゃりん"
	phraseLatinAmericanSpanish string = "chanchi"
	phraseKorean               string = "땡그랑"
	phraseRussian              string = "хрюк"
	phraseSpanish              string = "chanchi"
	phraseSimplifiedChinese    string = "当啷"
	phraseTraditionalChinese   string = "噹啷"
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
