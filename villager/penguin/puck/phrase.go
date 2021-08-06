package puck

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "brrrrrrrrr"
	phraseCanadianFrench       string = "brrrrrrrrr"
	phraseDutch                string = "brrrrrrrrr"
	phraseFrench               string = "brrrrrrrrr"
	phraseGerman               string = "brrrrrrrrr"
	phraseItalian              string = "brrrrrrrrr"
	phraseJapanese             string = "さぶー"
	phraseLatinAmericanSpanish string = "brrr-uah"
	phraseKorean               string = "엣취"
	phraseRussian              string = "брр"
	phraseSpanish              string = "sardinilla"
	phraseSimplifiedChinese    string = "候补"
	phraseTraditionalChinese   string = "候補"
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
