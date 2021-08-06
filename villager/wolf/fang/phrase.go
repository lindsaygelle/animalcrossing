package fang

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cha-chomp"
	phraseCanadianFrench       string = "sagouin"
	phraseDutch                string = "knauw"
	phraseFrench               string = "sagouin"
	phraseGerman               string = "braul"
	phraseItalian              string = "fabuuula"
	phraseJapanese             string = "ですダス"
	phraseLatinAmericanSpanish string = "en-ñac"
	phraseKorean               string = "콜록콜록"
	phraseRussian              string = "грызь"
	phraseSpanish              string = "en-ñac"
	phraseSimplifiedChinese    string = "这倒是"
	phraseTraditionalChinese   string = "這倒是"
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
