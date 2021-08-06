package flip

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "rerack"
	phraseCanadianFrench       string = "oh hisse"
	phraseDutch                string = "oe-⁠a-⁠a"
	phraseFrench               string = "oh hisse"
	phraseGerman               string = "kikiki"
	phraseItalian              string = "clap clap"
	phraseJapanese             string = "どっこい"
	phraseLatinAmericanSpanish string = "afuá"
	phraseKorean               string = "빠샤"
	phraseRussian              string = "у-у-у"
	phraseSpanish              string = "afuá"
	phraseSimplifiedChinese    string = "差不多"
	phraseTraditionalChinese   string = "差不多"
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
