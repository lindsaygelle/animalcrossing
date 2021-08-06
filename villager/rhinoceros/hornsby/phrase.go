package hornsby

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "schnozzle"
	phraseCanadianFrench       string = "euh"
	phraseDutch                string = "sneuzel"
	phraseFrench               string = "euh"
	phraseGerman               string = "schneuzel"
	phraseItalian              string = "splush"
	phraseJapanese             string = "のネン"
	phraseLatinAmericanSpanish string = "kaplut"
	phraseKorean               string = "뿔뿔"
	phraseRussian              string = "шнобель"
	phraseSpanish              string = "narigón"
	phraseSimplifiedChinese    string = "念念"
	phraseTraditionalChinese   string = "念念"
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
