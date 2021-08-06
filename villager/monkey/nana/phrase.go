package nana

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "po po"
	phraseCanadianFrench       string = "po po"
	phraseDutch                string = "banaan"
	phraseFrench               string = "po po"
	phraseGerman               string = "spitzi"
	phraseItalian              string = "po po"
	phraseJapanese             string = "ウキャ"
	phraseLatinAmericanSpanish string = "ukiú"
	phraseKorean               string = "끼끼"
	phraseRussian              string = "пу-пу"
	phraseSpanish              string = "monada"
	phraseSimplifiedChinese    string = "唔唔"
	phraseTraditionalChinese   string = "唔唔"
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
