package claude

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hopalong"
	phraseCanadianFrench       string = "sans rire"
	phraseDutch                string = "hopsala"
	phraseFrench               string = "sans rire"
	phraseGerman               string = "hüpfauf"
	phraseItalian              string = "hoppela"
	phraseJapanese             string = "ぶいぶい"
	phraseLatinAmericanSpanish string = "hópala"
	phraseKorean               string = "아으셔"
	phraseRussian              string = "скок-поскок"
	phraseSpanish              string = "hópala"
	phraseSimplifiedChinese    string = "酸酸"
	phraseTraditionalChinese   string = "酸酸"
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
