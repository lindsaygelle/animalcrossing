package tasha

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nice nice"
	phraseCanadianFrench       string = "ptigri"
	phraseDutch                string = "jaja"
	phraseFrench               string = "ptigri"
	phraseGerman               string = "tjaja"
	phraseItalian              string = "munch"
	phraseJapanese             string = "やるわね"
	phraseLatinAmericanSpanish string = "bi-buá"
	phraseKorean               string = "굿굿"
	phraseRussian              string = "ладно"
	phraseSpanish              string = "pincel"
	phraseSimplifiedChinese    string = "做得好"
	phraseTraditionalChinese   string = "做得好"
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
