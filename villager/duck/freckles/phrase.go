package freckles

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ducky"
	phraseCanadianFrench       string = "poussin"
	phraseDutch                string = "eendje"
	phraseFrench               string = "poussin"
	phraseGerman               string = "schnatter"
	phraseItalian              string = "qua qua"
	phraseJapanese             string = "ぎょぎょ"
	phraseLatinAmericanSpanish string = "cui-cuá"
	phraseKorean               string = "물물"
	phraseRussian              string = "утенок"
	phraseSpanish              string = "cui-cuá"
	phraseSimplifiedChinese    string = "鱼鱼"
	phraseTraditionalChinese   string = "魚魚"
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
