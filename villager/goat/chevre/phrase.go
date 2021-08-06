package chevre

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "la baa"
	phraseCanadianFrench       string = "la baa"
	phraseDutch                string = "la mèè"
	phraseFrench               string = "la baa"
	phraseGerman               string = "määääh"
	phraseItalian              string = "beeello"
	phraseJapanese             string = "っぺ"
	phraseLatinAmericanSpanish string = "beee-beee"
	phraseKorean               string = "맞아유"
	phraseRussian              string = "ле ме-е-е"
	phraseSpanish              string = "beee-beee"
	phraseSimplifiedChinese    string = "呗"
	phraseTraditionalChinese   string = "唄"
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
