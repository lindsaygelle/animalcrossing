package beardo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "whiskers"
	phraseCanadianFrench       string = "ciao"
	phraseDutch                string = "snorhaar"
	phraseFrench               string = "poilodos"
	phraseGerman               string = "brrroho"
	phraseItalian              string = "andale"
	phraseJapanese             string = "オッホン"
	phraseLatinAmericanSpanish string = "bigotre"
	phraseKorean               string = "오홍홍"
	phraseRussian              string = "усы мои"
	phraseSpanish              string = "bigotre"
	phraseSimplifiedChinese    string = "咳咳"
	phraseTraditionalChinese   string = "咳咳"
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
