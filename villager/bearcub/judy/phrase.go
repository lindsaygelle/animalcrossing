package judy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "myohmy"
	phraseCanadianFrench       string = "oh là là"
	phraseDutch                string = "nee maar"
	phraseFrench               string = "oh là là"
	phraseGerman               string = "träum"
	phraseItalian              string = "misu misu"
	phraseJapanese             string = "あらら"
	phraseLatinAmericanSpanish string = "uyuyuy"
	phraseKorean               string = "어머머"
	phraseRussian              string = "охохонюшки"
	phraseSpanish              string = "qué cosas"
	phraseSimplifiedChinese    string = "哎呀"
	phraseTraditionalChinese   string = "哎呀"
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
