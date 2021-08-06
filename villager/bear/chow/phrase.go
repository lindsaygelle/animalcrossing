package chow

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "aiya"
	phraseCanadianFrench       string = "lala"
	phraseDutch                string = "kiai"
	phraseFrench               string = "lala"
	phraseGerman               string = "hiijaa"
	phraseItalian              string = "ehilà"
	phraseJapanese             string = "アルヨ"
	phraseLatinAmericanSpanish string = "grinchu"
	phraseKorean               string = "그럼"
	phraseRussian              string = "ай-яй"
	phraseSpanish              string = "ándale-oso"
	phraseSimplifiedChinese    string = "有喔"
	phraseTraditionalChinese   string = "有喔"
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
