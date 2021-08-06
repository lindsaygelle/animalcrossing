package rilla

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hello"
	phraseCanadianFrench       string = "rorille"
	phraseDutch                string = "hello"
	phraseFrench               string = "rorille"
	phraseGerman               string = "kittykong"
	phraseItalian              string = "arrruuu"
	phraseJapanese             string = "ハロー"
	phraseLatinAmericanSpanish string = "kittybún"
	phraseKorean               string = "헬로"
	phraseRussian              string = "алло"
	phraseSpanish              string = "kittybún"
	phraseSimplifiedChinese    string = "哈罗"
	phraseTraditionalChinese   string = "哈囉"
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
