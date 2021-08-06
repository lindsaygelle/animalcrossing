package mira

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cottontail"
	phraseCanadianFrench       string = "zip zip"
	phraseDutch                string = "staartje"
	phraseFrench               string = "auch auch"
	phraseGerman               string = "jahopp"
	phraseItalian              string = "baffoplà"
	phraseJapanese             string = "だムーン"
	phraseLatinAmericanSpanish string = "hep-hep"
	phraseKorean               string = "거봐라"
	phraseRussian              string = "зайцехвост"
	phraseSpanish              string = "muesli"
	phraseSimplifiedChinese    string = "月球"
	phraseTraditionalChinese   string = "月球"
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
