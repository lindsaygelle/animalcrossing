package opal

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snoot"
	phraseCanadianFrench       string = "gropif"
	phraseDutch                string = "snoet"
	phraseFrench               string = "gropif"
	phraseGerman               string = "tröööt"
	phraseItalian              string = "snoob"
	phraseJapanese             string = "ヨン"
	phraseLatinAmericanSpanish string = "flip-flop"
	phraseKorean               string = "땡"
	phraseRussian              string = "хобот"
	phraseSpanish              string = "flip-flop"
	phraseSimplifiedChinese    string = "勇"
	phraseTraditionalChinese   string = "勇"
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
