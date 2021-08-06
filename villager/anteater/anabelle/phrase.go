package anabelle

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snorty"
	phraseCanadianFrench       string = "grognon"
	phraseDutch                string = "snork"
	phraseFrench               string = "grognon"
	phraseGerman               string = "puuust"
	phraseItalian              string = "snorty"
	phraseJapanese             string = "マジでー"
	phraseLatinAmericanSpanish string = "fa-fiú"
	phraseKorean               string = "정말"
	phraseRussian              string = "хрум-хрум"
	phraseSpanish              string = "esnoink"
	phraseSimplifiedChinese    string = "真的假的"
	phraseTraditionalChinese   string = "真的假的"
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
