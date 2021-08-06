package scoot

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "zip zoom"
	phraseCanadianFrench       string = "zak zak"
	phraseDutch                string = "broem"
	phraseFrench               string = "zak zak"
	phraseGerman               string = "flitzer"
	phraseItalian              string = "zip zoom"
	phraseJapanese             string = "グワッ"
	phraseLatinAmericanSpanish string = "zapizum"
	phraseKorean               string = "꾸왁"
	phraseRussian              string = "брум-брум"
	phraseSpanish              string = "zapizum"
	phraseSimplifiedChinese    string = "呱呱"
	phraseTraditionalChinese   string = "呱呱"
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
