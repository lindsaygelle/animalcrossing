package paolo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pal"
	phraseCanadianFrench       string = "fanfan"
	phraseDutch                string = "gabber"
	phraseFrench               string = "fanfan"
	phraseGerman               string = "pruuust"
	phraseItalian              string = "pistaaaa"
	phraseJapanese             string = "パオ"
	phraseLatinAmericanSpanish string = "pruuuf"
	phraseKorean               string = "파오"
	phraseRussian              string = "друг"
	phraseSpanish              string = "amigante"
	phraseSimplifiedChinese    string = "保罗"
	phraseTraditionalChinese   string = "保羅"
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
