package elvis

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "unh-hunh"
	phraseCanadianFrench       string = "bébé"
	phraseDutch                string = "aloha"
	phraseFrench               string = "bébé"
	phraseGerman               string = "grolll"
	phraseItalian              string = "unh-hunh"
	phraseJapanese             string = "ダロガ"
	phraseLatinAmericanSpanish string = "groar"
	phraseKorean               string = "안그냐"
	phraseRussian              string = "буги-вуги"
	phraseSpanish              string = "groar"
	phraseSimplifiedChinese    string = "听懂吧"
	phraseTraditionalChinese   string = "聽懂吧"
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
