package kabuki

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "meooo-OH"
	phraseCanadianFrench       string = "mia-ouh"
	phraseDutch                string = "mie-auw"
	phraseFrench               string = "mia-ouh"
	phraseGerman               string = "miiiau"
	phraseItalian              string = "meooo-OH"
	phraseJapanese             string = "ぃよぉー"
	phraseLatinAmericanSpanish string = "imikimono"
	phraseKorean               string = "꺄"
	phraseRussian              string = "мя-я-яу"
	phraseSpanish              string = "imikimono"
	phraseSimplifiedChinese    string = "咦唷"
	phraseTraditionalChinese   string = "咦唷"
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
