package ankha

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "me meow"
	phraseCanadianFrench       string = "mi-i-i-ouh"
	phraseDutch                string = "sfinx"
	phraseFrench               string = "mi-i-i-ouh"
	phraseGerman               string = "mi-miau"
	phraseItalian              string = "miao miao"
	phraseJapanese             string = "クフフ"
	phraseLatinAmericanSpanish string = "marramius"
	phraseKorean               string = "파트라"
	phraseRussian              string = "ми-мяу"
	phraseSpanish              string = "marramius"
	phraseSimplifiedChinese    string = "尼罗"
	phraseTraditionalChinese   string = "尼羅"
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
