package midge

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tweedledee"
	phraseCanadianFrench       string = "loulou"
	phraseDutch                string = "kwetter"
	phraseFrench               string = "loulou"
	phraseGerman               string = "twitiriti"
	phraseItalian              string = "firulì"
	phraseJapanese             string = "キョン"
	phraseLatinAmericanSpanish string = "arrurrú"
	phraseKorean               string = "핑그르르"
	phraseRussian              string = "тви-тви"
	phraseSpanish              string = "arrurrú"
	phraseSimplifiedChinese    string = "脸红"
	phraseTraditionalChinese   string = "臉紅"
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
