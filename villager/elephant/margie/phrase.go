package margie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tootie"
	phraseCanadianFrench       string = "toutie"
	phraseDutch                string = "pwèèpie"
	phraseFrench               string = "toutie"
	phraseGerman               string = "trampl"
	phraseItalian              string = "tootie"
	phraseJapanese             string = "シャララ"
	phraseLatinAmericanSpanish string = "fiuf-fiuf"
	phraseKorean               string = "샬랄라"
	phraseRussian              string = "тру-у-ти"
	phraseSpanish              string = "flaqui"
	phraseSimplifiedChinese    string = "莎啦啦"
	phraseTraditionalChinese   string = "莎啦啦"
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
