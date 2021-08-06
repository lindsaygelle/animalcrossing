package louie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hoo hoo ha"
	phraseCanadianFrench       string = "tonneau"
	phraseDutch                string = "oe-oe-ah"
	phraseFrench               string = "tonneau"
	phraseGerman               string = "ugh-ugh"
	phraseItalian              string = "bum bum"
	phraseJapanese             string = "コング"
	phraseLatinAmericanSpanish string = "gorilérico"
	phraseKorean               string = "하압"
	phraseRussian              string = "ух-ух-ах"
	phraseSpanish              string = "gorilérico"
	phraseSimplifiedChinese    string = "钢钢"
	phraseTraditionalChinese   string = "鋼鋼"
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
