package teddy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "grooof"
	phraseCanadianFrench       string = "grrrrrr"
	phraseDutch                string = "grommm"
	phraseFrench               string = "grrrrrr"
	phraseGerman               string = "gruff"
	phraseItalian              string = "gruf"
	phraseJapanese             string = "ですたい"
	phraseLatinAmericanSpanish string = "gruuuf"
	phraseKorean               string = "입네다"
	phraseRussian              string = "быр-быр"
	phraseSpanish              string = "gruuuf"
	phraseSimplifiedChinese    string = "太好了"
	phraseTraditionalChinese   string = "太好了"
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
