package cally

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "WHEE"
	phraseCanadianFrench       string = "cacahuète"
	phraseDutch                string = "cashew"
	phraseFrench               string = "yeeeees"
	phraseGerman               string = "huiii"
	phraseItalian              string = "SCRUNCH"
	phraseJapanese             string = "ララー"
	phraseLatinAmericanSpanish string = "ñip-ñip"
	phraseKorean               string = "랄랄라"
	phraseRussian              string = "ух ты"
	phraseSpanish              string = "macadamia"
	phraseSimplifiedChinese    string = "啦啦"
	phraseTraditionalChinese   string = "啦啦"
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
