package lionel

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "precisely"
	phraseCanadianFrench       string = "papatte"
	phraseDutch                string = "exact"
	phraseFrench               string = "papatte"
	phraseGerman               string = "grrroooh"
	phraseItalian              string = "ahora"
	phraseJapanese             string = "まさしく"
	phraseLatinAmericanSpanish string = "picoplás"
	phraseKorean               string = "옳거니"
	phraseRussian              string = "честь имею"
	phraseSpanish              string = "picoplás"
	phraseSimplifiedChinese    string = "正确"
	phraseTraditionalChinese   string = "正確"
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
