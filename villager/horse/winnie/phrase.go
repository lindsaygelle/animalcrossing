package winnie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hay-OK"
	phraseCanadianFrench       string = "youp là"
	phraseDutch                string = "knol"
	phraseFrench               string = "youp là"
	phraseGerman               string = "heu-ja"
	phraseItalian              string = "biadina"
	phraseJapanese             string = "みゃあ"
	phraseLatinAmericanSpanish string = "okeimakei"
	phraseKorean               string = "당근당근"
	phraseRussian              string = "тык-дык"
	phraseSpanish              string = "okeimakei"
	phraseSimplifiedChinese    string = "嘶嘶"
	phraseTraditionalChinese   string = "嘶嘶"
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
