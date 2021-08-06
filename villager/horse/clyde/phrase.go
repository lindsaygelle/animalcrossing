package clyde

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "clip clawp"
	phraseCanadianFrench       string = "mouh-ouh"
	phraseDutch                string = "klipklop"
	phraseFrench               string = "mouh-ouh"
	phraseGerman               string = "schauder"
	phraseItalian              string = "clop clop"
	phraseJapanese             string = "ぴろぴろ"
	phraseLatinAmericanSpanish string = "clip-clop"
	phraseKorean               string = "다그닥"
	phraseRussian              string = "цок-цок"
	phraseSpanish              string = "troteras"
	phraseSimplifiedChinese    string = "吹卷卷"
	phraseTraditionalChinese   string = "吹捲捲"
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
