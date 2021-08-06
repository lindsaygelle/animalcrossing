package apollo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pah"
	phraseCanadianFrench       string = "trraaaa"
	phraseDutch                string = "poeh"
	phraseFrench               string = "trraaaa"
	phraseGerman               string = "ahhh"
	phraseItalian              string = "pah"
	phraseJapanese             string = "だワイ"
	phraseLatinAmericanSpanish string = "rapahhh"
	phraseKorean               string = "다꿇어"
	phraseRussian              string = "пф-ф"
	phraseSpanish              string = "rapaz"
	phraseSimplifiedChinese    string = "唔咿"
	phraseTraditionalChinese   string = "唔咿"
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
