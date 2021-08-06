package rolf

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "grrrolf"
	phraseCanadianFrench       string = "grrrufff"
	phraseDutch                string = "grrrolf"
	phraseFrench               string = "grrrufff"
	phraseGerman               string = "brüll"
	phraseItalian              string = "grrrolf"
	phraseJapanese             string = "だガー"
	phraseLatinAmericanSpanish string = "grooorf"
	phraseKorean               string = "근데"
	phraseRussian              string = "гр-р-рольф"
	phraseSpanish              string = "grooorf"
	phraseSimplifiedChinese    string = "王"
	phraseTraditionalChinese   string = "王"
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
