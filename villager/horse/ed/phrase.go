package ed

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "greenhorn"
	phraseCanadianFrench       string = "tagada"
	phraseDutch                string = "ju"
	phraseFrench               string = "tagada"
	phraseGerman               string = "greenhorn"
	phraseItalian              string = "pirulì"
	phraseJapanese             string = "じゃない"
	phraseLatinAmericanSpanish string = "jo-joy"
	phraseKorean               string = "게슴츠레"
	phraseRussian              string = "новичок"
	phraseSpanish              string = "jo-joy"
	phraseSimplifiedChinese    string = "不是啦"
	phraseTraditionalChinese   string = "不是啦"
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
