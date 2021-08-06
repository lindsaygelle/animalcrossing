package merry

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mweee"
	phraseCanadianFrench       string = "miou-mi"
	phraseDutch                string = "mi-jippie"
	phraseFrench               string = "miou-mi"
	phraseGerman               string = "prrrrr"
	phraseItalian              string = "pffffft"
	phraseJapanese             string = "にゃん"
	phraseLatinAmericanSpanish string = "michifú"
	phraseKorean               string = "냥냥"
	phraseRussian              string = "мяуи-и-и"
	phraseSpanish              string = "michifú"
	phraseSimplifiedChinese    string = "喵嗯"
	phraseTraditionalChinese   string = "喵嗯"
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
