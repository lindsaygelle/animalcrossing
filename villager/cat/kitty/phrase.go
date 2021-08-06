package kitty

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mrowrr"
	phraseCanadianFrench       string = "miaaaouh"
	phraseDutch                string = "giechel"
	phraseFrench               string = "miaaaouh"
	phraseGerman               string = "miezmiez"
	phraseItalian              string = "fusolo"
	phraseJapanese             string = "フフ"
	phraseLatinAmericanSpanish string = "miaul"
	phraseKorean               string = "자갸"
	phraseRussian              string = "мр-мр-мр"
	phraseSpanish              string = "cascabel"
	phraseSimplifiedChinese    string = "古古"
	phraseTraditionalChinese   string = "古古"
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
