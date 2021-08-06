package bitty

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "my dear"
	phraseCanadianFrench       string = "joli cœur"
	phraseDutch                string = "liefje"
	phraseFrench               string = "joli cœur"
	phraseGerman               string = "mei o mei"
	phraseItalian              string = "bibù"
	phraseJapanese             string = "だピー"
	phraseLatinAmericanSpanish string = "hipopop"
	phraseKorean               string = "쩝쩝"
	phraseRussian              string = "дорогуша"
	phraseSpanish              string = "corassón"
	phraseSimplifiedChinese    string = "哔哔"
	phraseTraditionalChinese   string = "嗶嗶"
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
