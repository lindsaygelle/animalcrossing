package rex

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cool cat"
	phraseCanadianFrench       string = "lionceau"
	phraseDutch                string = "coole kat"
	phraseFrench               string = "lionceau"
	phraseGerman               string = "coolcat"
	phraseItalian              string = "grrringo"
	phraseJapanese             string = "だオン"
	phraseLatinAmericanSpanish string = "grrreh"
	phraseKorean               string = "티라이온"
	phraseRussian              string = "спокойно"
	phraseSpanish              string = "melenas"
	phraseSimplifiedChinese    string = "狮狮"
	phraseTraditionalChinese   string = "獅獅"
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
