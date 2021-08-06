package savannah

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "y'all"
	phraseCanadianFrench       string = "wimbowé"
	phraseDutch                string = "pony"
	phraseFrench               string = "wimbowé"
	phraseGerman               string = "wichichi"
	phraseItalian              string = "hip hip"
	phraseJapanese             string = "ってば"
	phraseLatinAmericanSpanish string = "azuquita"
	phraseKorean               string = "맞아요"
	phraseRussian              string = "и все"
	phraseSpanish              string = "azucarillo"
	phraseSimplifiedChinese    string = "说到"
	phraseTraditionalChinese   string = "說到"
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
