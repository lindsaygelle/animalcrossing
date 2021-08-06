package rhonda

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bigfoot"
	phraseCanadianFrench       string = "grand pied"
	phraseDutch                string = "grootvoet"
	phraseFrench               string = "grand pied"
	phraseGerman               string = "stampf"
	phraseItalian              string = "piedone"
	phraseJapanese             string = "ンフ"
	phraseLatinAmericanSpanish string = "jip-jip"
	phraseKorean               string = "훗"
	phraseRussian              string = "босс"
	phraseSpanish              string = "jip-jip"
	phraseSimplifiedChinese    string = "嗯噗"
	phraseTraditionalChinese   string = "嗯噗"
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
