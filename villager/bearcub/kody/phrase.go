package kody

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "grah grah"
	phraseCanadianFrench       string = "grih grih"
	phraseDutch                string = "okidoki"
	phraseFrench               string = "grih grih"
	phraseGerman               string = "okeydokey"
	phraseItalian              string = "oki-doki"
	phraseJapanese             string = "のこころ"
	phraseLatinAmericanSpanish string = "graaau"
	phraseKorean               string = "나는나야"
	phraseRussian              string = "гррр"
	phraseSpanish              string = "mola"
	phraseSimplifiedChinese    string = "精神"
	phraseTraditionalChinese   string = "精神"
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
