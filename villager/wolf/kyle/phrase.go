package kyle

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "alpha"
	phraseCanadianFrench       string = "garooouuu"
	phraseDutch                string = "alfa"
	phraseFrench               string = "garooouuu"
	phraseGerman               string = "ahuuu"
	phraseItalian              string = "sgnack"
	phraseJapanese             string = "オゥイェ"
	phraseLatinAmericanSpanish string = "aujujú"
	phraseKorean               string = "오우예"
	phraseRussian              string = "вожак"
	phraseSpanish              string = "aujujú"
	phraseSimplifiedChinese    string = "喔耶"
	phraseTraditionalChinese   string = "喔耶"
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
