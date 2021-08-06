package cranston

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sweatband"
	phraseCanadianFrench       string = "omelette"
	phraseDutch                string = "zweetband"
	phraseFrench               string = "omelette"
	phraseGerman               string = "happahappa"
	phraseItalian              string = "boa"
	phraseJapanese             string = "およよ"
	phraseLatinAmericanSpanish string = "gruqui"
	phraseKorean               string = "오요용"
	phraseRussian              string = "не потей"
	phraseSpanish              string = "gruqui"
	phraseSimplifiedChinese    string = "哎呀呀"
	phraseTraditionalChinese   string = "哎呀呀"
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
