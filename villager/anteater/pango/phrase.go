package pango

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snooooof"
	phraseCanadianFrench       string = "pouuuuuf"
	phraseDutch                string = "snufffff"
	phraseFrench               string = "pouuuuuf"
	phraseGerman               string = "schnieef"
	phraseItalian              string = "snuuf"
	phraseJapanese             string = "だっしー"
	phraseLatinAmericanSpanish string = "snuf-snuf"
	phraseKorean               string = "라지요"
	phraseRussian              string = "снуф-снуф"
	phraseSpanish              string = "snuf-snuf"
	phraseSimplifiedChinese    string = "希希"
	phraseTraditionalChinese   string = "希希"
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
