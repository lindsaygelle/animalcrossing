package daisy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bow-WOW"
	phraseCanadianFrench       string = "mon chou"
	phraseDutch                string = "hazewind"
	phraseFrench               string = "mon chou"
	phraseGerman               string = "wauwau"
	phraseItalian              string = "bau WOW"
	phraseJapanese             string = "だよね"
	phraseLatinAmericanSpanish string = "guaucito"
	phraseKorean               string = "그렇죠"
	phraseRussian              string = "тяв-ого"
	phraseSpanish              string = "guaay"
	phraseSimplifiedChinese    string = "对不对"
	phraseTraditionalChinese   string = "對不對"
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
