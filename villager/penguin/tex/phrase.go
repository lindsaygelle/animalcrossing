package tex

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "picante"
	phraseCanadianFrench       string = "gla gla"
	phraseDutch                string = "picante"
	phraseFrench               string = "gla gla"
	phraseGerman               string = "watschl"
	phraseItalian              string = "frio"
	phraseJapanese             string = "ベイベッ"
	phraseLatinAmericanSpanish string = "siclaro"
	phraseKorean               string = "베이베"
	phraseRussian              string = "огонек"
	phraseSpanish              string = "siclaro"
	phraseSimplifiedChinese    string = "宝宝"
	phraseTraditionalChinese   string = "寶寶"
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
