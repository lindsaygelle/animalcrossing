package soleil

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tarnation"
	phraseCanadianFrench       string = "ciel"
	phraseDutch                string = "korenwolf"
	phraseFrench               string = "gnognotte"
	phraseGerman               string = "hümpf"
	phraseItalian              string = "vriiin"
	phraseJapanese             string = "だってば"
	phraseLatinAmericanSpanish string = "ñiqui-ñú"
	phraseKorean               string = "파샤샤"
	phraseRussian              string = "о-хо-хо"
	phraseSpanish              string = "moflete"
	phraseSimplifiedChinese    string = "相信我"
	phraseTraditionalChinese   string = "相信我"
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
