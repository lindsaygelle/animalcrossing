package ribbot

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "zzrrbbit"
	phraseCanadianFrench       string = "croac"
	phraseDutch                string = "kwwwrrrk"
	phraseFrench               string = "croac"
	phraseGerman               string = "quarrrk"
	phraseItalian              string = "cracrà"
	phraseJapanese             string = "だロボ"
	phraseLatinAmericanSpanish string = "crobit"
	phraseKorean               string = "오버"
	phraseRussian              string = "квак-дзынь"
	phraseSpanish              string = "crobit"
	phraseSimplifiedChinese    string = "机器"
	phraseTraditionalChinese   string = "機器"
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
