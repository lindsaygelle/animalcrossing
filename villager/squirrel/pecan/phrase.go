package pecan

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chipmunk"
	phraseCanadianFrench       string = "fouloulou"
	phraseDutch                string = "eekhoorn"
	phraseFrench               string = "fouloulou"
	phraseGerman               string = "spatzl"
	phraseItalian              string = "cipì"
	phraseJapanese             string = "つんっ"
	phraseLatinAmericanSpanish string = "piñoní"
	phraseKorean               string = "쯧쯧"
	phraseRussian              string = "бурундук"
	phraseSpanish              string = "piñoncito"
	phraseSimplifiedChinese    string = "正经"
	phraseTraditionalChinese   string = "正經"
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
