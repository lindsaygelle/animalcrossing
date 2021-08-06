package genji

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mochi"
	phraseCanadianFrench       string = "mochi"
	phraseDutch                string = "mochi"
	phraseFrench               string = "mochi"
	phraseGerman               string = "goseimas"
	phraseItalian              string = "mochi"
	phraseJapanese             string = "まろ"
	phraseLatinAmericanSpanish string = "mochi"
	phraseKorean               string = "땡글"
	phraseRussian              string = "зайпончик"
	phraseSpanish              string = "mochi"
	phraseSimplifiedChinese    string = "臣"
	phraseTraditionalChinese   string = "臣"
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
