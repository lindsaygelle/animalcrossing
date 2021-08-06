package anicotti

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cannoli"
	phraseCanadianFrench       string = "dis donc"
	phraseDutch                string = "cannoli"
	phraseFrench               string = "dis donc"
	phraseGerman               string = "piepser"
	phraseItalian              string = "cippoli"
	phraseJapanese             string = "ルンルン"
	phraseLatinAmericanSpanish string = "cloricló"
	phraseKorean               string = "룰룰"
	phraseRussian              string = "канноли"
	phraseSpanish              string = "cloricló"
	phraseSimplifiedChinese    string = "开心"
	phraseTraditionalChinese   string = "開心"
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
