package flurry

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "powderpuff"
	phraseCanadianFrench       string = "abajoujou"
	phraseDutch                string = "donsbal"
	phraseFrench               string = "abajoujou"
	phraseGerman               string = "millimilli"
	phraseItalian              string = "fondue"
	phraseJapanese             string = "なのです"
	phraseLatinAmericanSpanish string = "nomnom"
	phraseKorean               string = "뽀드득"
	phraseRussian              string = "пушистик"
	phraseSpanish              string = "nomnom"
	phraseSimplifiedChinese    string = "对啊"
	phraseTraditionalChinese   string = "對啊"
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
