package blanche

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "quite so"
	phraseCanadianFrench       string = "viiiite"
	phraseDutch                string = "juist ja"
	phraseFrench               string = "viiiite"
	phraseGerman               string = "flaum"
	phraseItalian              string = "bien sur"
	phraseJapanese             string = "ですのね"
	phraseLatinAmericanSpanish string = "igh-igh"
	phraseKorean               string = "그랬구나"
	phraseRussian              string = "вот так"
	phraseSpanish              string = "goticas"
	phraseSimplifiedChinese    string = "是嘛"
	phraseTraditionalChinese   string = "是嘛"
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
