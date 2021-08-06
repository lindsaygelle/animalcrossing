package avery

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "skree-haw"
	phraseCanadianFrench       string = "grrraak"
	phraseDutch                string = "kra-hoe"
	phraseFrench               string = "phélès"
	phraseGerman               string = "flapp"
	phraseItalian              string = "gawk"
	phraseJapanese             string = "アリョイ"
	phraseLatinAmericanSpanish string = "grrraak"
	phraseKorean               string = "내놔"
	phraseRussian              string = "курлы"
	phraseSpanish              string = "grrraak"
	phraseSimplifiedChinese    string = "溜溜"
	phraseTraditionalChinese   string = "溜溜"
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
