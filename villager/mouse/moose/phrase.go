package moose

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "shorty"
	phraseCanadianFrench       string = "bééléé"
	phraseDutch                string = "kleintje"
	phraseFrench               string = "bééléé"
	phraseGerman               string = "käääse"
	phraseItalian              string = "quink"
	phraseJapanese             string = "にゃろ"
	phraseLatinAmericanSpanish string = "atiza"
	phraseKorean               string = "짜샤"
	phraseRussian              string = "коротышка"
	phraseSpanish              string = "atiza"
	phraseSimplifiedChinese    string = "可恶"
	phraseTraditionalChinese   string = "可惡"
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
