package whitney

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snappy"
	phraseCanadianFrench       string = "bing bang"
	phraseDutch                string = "hap"
	phraseFrench               string = "sbam"
	phraseGerman               string = "jaulll"
	phraseItalian              string = "snappy"
	phraseJapanese             string = "ステキね"
	phraseLatinAmericanSpanish string = "auf-auf"
	phraseKorean               string = "멋져"
	phraseRussian              string = "цап"
	phraseSpanish              string = "auf-auf"
	phraseSimplifiedChinese    string = "漂亮"
	phraseTraditionalChinese   string = "漂亮"
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
