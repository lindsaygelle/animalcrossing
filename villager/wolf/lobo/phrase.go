package lobo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ah-rooooo"
	phraseCanadianFrench       string = "ah rooooo"
	phraseDutch                string = "aooooeee"
	phraseFrench               string = "ah rooooo"
	phraseGerman               string = "auuuuuu"
	phraseItalian              string = "etchuuuu"
	phraseJapanese             string = "だぜよ"
	phraseLatinAmericanSpanish string = "auuuh"
	phraseKorean               string = "봐봐"
	phraseRussian              string = "ау-у-у"
	phraseSpanish              string = "auuuh"
	phraseSimplifiedChinese    string = "不然呢"
	phraseTraditionalChinese   string = "不然呢"
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
