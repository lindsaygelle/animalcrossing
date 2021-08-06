package pietro

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "honk honk"
	phraseCanadianFrench       string = "hahaha"
	phraseDutch                string = "alaaf"
	phraseFrench               string = "bamboche"
	phraseGerman               string = "helau"
	phraseItalian              string = "peco peco"
	phraseJapanese             string = "グフフ"
	phraseLatinAmericanSpanish string = "moquimoqui"
	phraseKorean               string = "므흐흐"
	phraseRussian              string = "бе-бе"
	phraseSpanish              string = "moquimoqui"
	phraseSimplifiedChinese    string = "丑丑"
	phraseTraditionalChinese   string = "丑丑"
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
