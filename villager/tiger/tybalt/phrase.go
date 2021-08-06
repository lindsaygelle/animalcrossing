package tybalt

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "grrrRAH"
	phraseCanadianFrench       string = "grouaaah"
	phraseDutch                string = "brrrrulll"
	phraseFrench               string = "grouaaah"
	phraseGerman               string = "grrrummel"
	phraseItalian              string = "grrrRAH"
	phraseJapanese             string = "だトラ"
	phraseLatinAmericanSpanish string = "raaarruf"
	phraseKorean               string = "으르렁"
	phraseRussian              string = "р-рык"
	phraseSpanish              string = "bengala"
	phraseSimplifiedChinese    string = "虎虎"
	phraseTraditionalChinese   string = "虎虎"
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
