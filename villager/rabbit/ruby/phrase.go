package ruby

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "li'l ears"
	phraseCanadianFrench       string = "n'oreille"
	phraseDutch                string = "hangoor"
	phraseFrench               string = "n'oreille"
	phraseGerman               string = "paffpaff"
	phraseItalian              string = "orecchine"
	phraseJapanese             string = "フツーに"
	phraseLatinAmericanSpanish string = "orejí"
	phraseKorean               string = "보통이지"
	phraseRussian              string = "ушки"
	phraseSpanish              string = "orejitas"
	phraseSimplifiedChinese    string = "普通"
	phraseTraditionalChinese   string = "普通"
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
