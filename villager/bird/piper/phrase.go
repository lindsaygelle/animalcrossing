package piper

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chickadee"
	phraseCanadianFrench       string = "puic puic"
	phraseDutch                string = "tjiep tjiep"
	phraseFrench               string = "puic puic"
	phraseGerman               string = "tschikadii"
	phraseItalian              string = "cocoricò"
	phraseJapanese             string = "けどー"
	phraseLatinAmericanSpanish string = "chupiyá"
	phraseKorean               string = "허나"
	phraseRussian              string = "чики-чик"
	phraseSpanish              string = "chupiyá"
	phraseSimplifiedChinese    string = "山雀"
	phraseTraditionalChinese   string = "山雀"
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
