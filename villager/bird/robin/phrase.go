package robin

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "la-di-da"
	phraseCanadianFrench       string = "ta-di-da"
	phraseDutch                string = "lalala"
	phraseFrench               string = "ta-di-da"
	phraseGerman               string = "ladida"
	phraseItalian              string = "yuppidù"
	phraseJapanese             string = "さ"
	phraseLatinAmericanSpanish string = "larará"
	phraseKorean               string = "글쎄"
	phraseRussian              string = "ля-ля-ля"
	phraseSpanish              string = "picatoste"
	phraseSimplifiedChinese    string = "喳"
	phraseTraditionalChinese   string = "喳"
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
