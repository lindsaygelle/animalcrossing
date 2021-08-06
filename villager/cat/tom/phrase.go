package tom

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "me-YOWZA"
	phraseCanadianFrench       string = "mistigri"
	phraseDutch                string = "miauw moe"
	phraseFrench               string = "mistigri"
	phraseGerman               string = "fauch"
	phraseItalian              string = "mi-IAO"
	phraseJapanese             string = "ナックル"
	phraseLatinAmericanSpanish string = "miooou"
	phraseKorean               string = "쳇"
	phraseRussian              string = "мя-я-яоза"
	phraseSpanish              string = "miooou"
	phraseSimplifiedChinese    string = "切"
	phraseTraditionalChinese   string = "關節"
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
