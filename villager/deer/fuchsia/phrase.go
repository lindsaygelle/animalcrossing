package fuchsia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "precious"
	phraseCanadianFrench       string = "frambou"
	phraseDutch                string = "meid"
	phraseFrench               string = "rosace"
	phraseGerman               string = "zick"
	phraseItalian              string = "cervilà"
	phraseJapanese             string = "なんしか"
	phraseLatinAmericanSpanish string = "braaaam"
	phraseKorean               string = "핑크크"
	phraseRussian              string = "зазноба"
	phraseSpanish              string = "braaaam"
	phraseSimplifiedChinese    string = "西卡"
	phraseTraditionalChinese   string = "西卡"
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
