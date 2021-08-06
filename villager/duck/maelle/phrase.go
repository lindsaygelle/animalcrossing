package maelle

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "duckling"
	phraseCanadianFrench       string = "canardou"
	phraseDutch                string = "pulletje"
	phraseFrench               string = "canardou"
	phraseGerman               string = "entchen"
	phraseItalian              string = "paperotto"
	phraseJapanese             string = "ふぅ"
	phraseLatinAmericanSpanish string = "plumi-piú"
	phraseKorean               string = "우아"
	phraseRussian              string = "желторотик"
	phraseSpanish              string = "perla"
	phraseSimplifiedChinese    string = "叹"
	phraseTraditionalChinese   string = "嘆"
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
