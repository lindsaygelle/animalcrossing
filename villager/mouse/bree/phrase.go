package bree

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cheeseball"
	phraseCanadianFrench       string = "boulette"
	phraseDutch                string = "kaaskoekje"
	phraseFrench               string = "boulette"
	phraseGerman               string = "kicher"
	phraseItalian              string = "formaggino"
	phraseJapanese             string = "なーんて"
	phraseLatinAmericanSpanish string = "quesuá"
	phraseKorean               string = "글치머"
	phraseRussian              string = "сы-ы-ыр"
	phraseSpanish              string = "orejas"
	phraseSimplifiedChinese    string = "说说的"
	phraseTraditionalChinese   string = "說說的"
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
