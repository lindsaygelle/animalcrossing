package tad

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sluuuurp"
	phraseCanadianFrench       string = "nénuf"
	phraseDutch                string = "sluuuurp"
	phraseFrench               string = "nénuf"
	phraseGerman               string = "schlürpf"
	phraseItalian              string = "sluuuurp"
	phraseJapanese             string = "だよん"
	phraseLatinAmericanSpanish string = "zump"
	phraseKorean               string = "흐압"
	phraseRussian              string = "жабррр"
	phraseSpanish              string = "zump"
	phraseSimplifiedChinese    string = "蝌蚪"
	phraseTraditionalChinese   string = "蝌蚪"
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
