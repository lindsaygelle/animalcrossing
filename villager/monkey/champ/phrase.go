package champ

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "choo CHOO"
	phraseCanadianFrench       string = "tchouTCHOU"
	phraseDutch                string = ""
	phraseFrench               string = "tchouTCHOU"
	phraseGerman               string = "u-u-u"
	phraseItalian              string = "choo CHOO"
	phraseJapanese             string = "ウッキー"
	phraseLatinAmericanSpanish string = "uh-ah-ah"
	phraseKorean               string = "야아아"
	phraseRussian              string = ""
	phraseSpanish              string = "uh-ah-ah"
	phraseSimplifiedChinese    string = ""
	phraseTraditionalChinese   string = ""
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
