package rocco

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hippie"
	phraseCanadianFrench       string = "babi"
	phraseDutch                string = "hippo"
	phraseFrench               string = "babi"
	phraseGerman               string = "pfffft"
	phraseItalian              string = "ipper"
	phraseJapanese             string = "だぎゃー"
	phraseLatinAmericanSpanish string = "hiponó"
	phraseKorean               string = "그러마"
	phraseRussian              string = "гиппи"
	phraseSpanish              string = "batracio"
	phraseSimplifiedChinese    string = "觉得是啦"
	phraseTraditionalChinese   string = "覺得是啦"
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
