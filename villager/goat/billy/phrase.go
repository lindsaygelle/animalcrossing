package billy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "dagnaabit"
	phraseCanadianFrench       string = "bêê quoi"
	phraseDutch                string = "nôh"
	phraseFrench               string = "eh bââlèze"
	phraseGerman               string = "o-o-ja-a-a"
	phraseItalian              string = "sorbeeelli"
	phraseJapanese             string = "よのう"
	phraseLatinAmericanSpanish string = "beerree"
	phraseKorean               string = "알아줘"
	phraseRussian              string = "такие-сякие"
	phraseSpanish              string = "fenómeno"
	phraseSimplifiedChinese    string = "唷喏"
	phraseTraditionalChinese   string = "唷喏"
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
