package wolfgang

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snarrrl"
	phraseCanadianFrench       string = "snurffff"
	phraseDutch                string = "snauw"
	phraseFrench               string = "snurffff"
	phraseGerman               string = "knurrr"
	phraseItalian              string = "zanna"
	phraseJapanese             string = "のな"
	phraseLatinAmericanSpanish string = "grauuuh"
	phraseKorean               string = "거봐"
	phraseRussian              string = "ау-рр"
	phraseSpanish              string = "grauuuh"
	phraseSimplifiedChinese    string = "罗罗"
	phraseTraditionalChinese   string = "羅羅"
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
