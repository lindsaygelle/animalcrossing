package drago

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "burrrn"
	phraseCanadianFrench       string = "ducroc"
	phraseDutch                string = "puf"
	phraseFrench               string = "ducroc"
	phraseGerman               string = "hömpf"
	phraseItalian              string = "tic tac"
	phraseJapanese             string = "えーと"
	phraseLatinAmericanSpanish string = "ñami"
	phraseKorean               string = "띠용띠용"
	phraseRussian              string = "огонь"
	phraseSpanish              string = "quemarrr"
	phraseSimplifiedChinese    string = "然后"
	phraseTraditionalChinese   string = "然後"
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
