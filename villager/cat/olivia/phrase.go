package olivia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "purrr"
	phraseCanadianFrench       string = "rrrrrrrrr"
	phraseDutch                string = "prrrrrr"
	phraseFrench               string = "rrrrrrrrr"
	phraseGerman               string = "schnurr"
	phraseItalian              string = "purrr"
	phraseJapanese             string = "なんやん"
	phraseLatinAmericanSpanish string = "michimiau"
	phraseKorean               string = "머냐"
	phraseRussian              string = "мур-мяу"
	phraseSpanish              string = "raspas"
	phraseSimplifiedChinese    string = "什么"
	phraseTraditionalChinese   string = "什麼"
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
