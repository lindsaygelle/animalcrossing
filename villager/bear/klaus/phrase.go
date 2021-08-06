package klaus

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "strudel"
	phraseCanadianFrench       string = "achtung"
	phraseDutch                string = "strudel"
	phraseFrench               string = "achtung"
	phraseGerman               string = "gruaaa"
	phraseItalian              string = "bailar"
	phraseJapanese             string = "オパー"
	phraseLatinAmericanSpanish string = "achurf"
	phraseKorean               string = "라저"
	phraseRussian              string = "штрудель"
	phraseSpanish              string = "ricitos"
	phraseSimplifiedChinese    string = "Over"
	phraseTraditionalChinese   string = "Over"
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
