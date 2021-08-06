package quillson

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ridukulous"
	phraseCanadianFrench       string = "arrêêêête"
	phraseDutch                string = "smient"
	phraseFrench               string = "arrêêêête"
	phraseGerman               string = "pluster"
	phraseItalian              string = "cua cua"
	phraseJapanese             string = "マジかよ"
	phraseLatinAmericanSpanish string = "cuacoac"
	phraseKorean               string = "리얼리"
	phraseRussian              string = "шу-утка"
	phraseSpanish              string = "verdoncho"
	phraseSimplifiedChinese    string = "说真的"
	phraseTraditionalChinese   string = "說真的"
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
