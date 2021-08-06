package sylvana

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hubbub"
	phraseCanadianFrench       string = "grignote"
	phraseDutch                string = "jammie"
	phraseFrench               string = "grignote"
	phraseGerman               string = "jammi"
	phraseItalian              string = "sgranocc"
	phraseJapanese             string = "ひゅん"
	phraseLatinAmericanSpanish string = "yupip"
	phraseKorean               string = "퓨우"
	phraseRussian              string = "хрум"
	phraseSpanish              string = "caldito"
	phraseSimplifiedChinese    string = "咻"
	phraseTraditionalChinese   string = "咻"
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
