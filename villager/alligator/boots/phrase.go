package boots

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "munchie"
	phraseCanadianFrench       string = "miam miam"
	phraseDutch                string = "smikkel"
	phraseFrench               string = "mariolle"
	phraseGerman               string = "schmatzi"
	phraseItalian              string = "gnammi"
	phraseJapanese             string = "だぴょん"
	phraseLatinAmericanSpanish string = "crococroco"
	phraseKorean               string = "만세"
	phraseRussian              string = "кусь"
	phraseSpanish              string = "crococroco"
	phraseSimplifiedChinese    string = "跳"
	phraseTraditionalChinese   string = "跳"
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
