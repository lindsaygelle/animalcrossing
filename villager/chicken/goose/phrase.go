package goose

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "buh-kay"
	phraseCanadianFrench       string = "bouyaka"
	phraseDutch                string = "toktok"
	phraseFrench               string = "bouyaka"
	phraseGerman               string = "buaa-ka-ka"
	phraseItalian              string = "cucù"
	phraseJapanese             string = "だコケ"
	phraseLatinAmericanSpanish string = "bukááá"
	phraseKorean               string = "키득"
	phraseRussian              string = "ко-кей"
	phraseSpanish              string = "caracol"
	phraseSimplifiedChinese    string = "咕咕"
	phraseTraditionalChinese   string = "咕咕"
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
