package baabara

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "daahling"
	phraseCanadianFrench       string = "chêêêri"
	phraseDutch                string = "schaaaapje"
	phraseFrench               string = "chêêêri"
	phraseGerman               string = "däääää"
	phraseItalian              string = "breeezza"
	phraseJapanese             string = "アンドゥ"
	phraseLatinAmericanSpanish string = "peluuusas"
	phraseKorean               string = "앤쥬"
	phraseRussian              string = "мила-ашка"
	phraseSpanish              string = "peeelusa"
	phraseSimplifiedChinese    string = "一二三"
	phraseTraditionalChinese   string = "一二三"
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
