package nan

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "kid"
	phraseCanadianFrench       string = "fromage"
	phraseDutch                string = "geitje"
	phraseFrench               string = "fromage"
	phraseGerman               string = "mekker"
	phraseItalian              string = "bimbi"
	phraseJapanese             string = "っしょ"
	phraseLatinAmericanSpanish string = "bibi-bee"
	phraseKorean               string = "그랬슈"
	phraseRussian              string = "козленок"
	phraseSpanish              string = "peque"
	phraseSimplifiedChinese    string = "书"
	phraseTraditionalChinese   string = "書"
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
