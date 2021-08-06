package egbert

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "doodle-duh"
	phraseCanadianFrench       string = "di di dou"
	phraseDutch                string = "kukelejus"
	phraseFrench               string = "di di dou"
	phraseGerman               string = "duudle-du"
	phraseItalian              string = "du du du"
	phraseJapanese             string = "だヨ"
	phraseLatinAmericanSpanish string = "kikirikí"
	phraseKorean               string = "짜잔"
	phraseRussian              string = "кукареку"
	phraseSpanish              string = "kikirikí"
	phraseSimplifiedChinese    string = "很优"
	phraseTraditionalChinese   string = "很優"
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
