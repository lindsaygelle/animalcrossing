package alli

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "graaagh"
	phraseCanadianFrench       string = "graaagh"
	phraseDutch                string = "graaach"
	phraseFrench               string = "graaagh"
	phraseGerman               string = "graaach"
	phraseItalian              string = "graaag"
	phraseJapanese             string = "どすえ"
	phraseLatinAmericanSpanish string = "muamú"
	phraseKorean               string = "얘야"
	phraseRussian              string = "гра-а-а"
	phraseSpanish              string = "lagartija"
	phraseSimplifiedChinese    string = "鳄鱼皮"
	phraseTraditionalChinese   string = "鱷魚皮"
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
