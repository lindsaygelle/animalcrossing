package bettina

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "eekers"
	phraseCanadianFrench       string = "skouik-ik"
	phraseDutch                string = "piepers"
	phraseFrench               string = "skouik-ik"
	phraseGerman               string = "achja"
	phraseItalian              string = "zink zink"
	phraseJapanese             string = "ですよね"
	phraseLatinAmericanSpanish string = "escui-cui"
	phraseKorean               string = "내말이요"
	phraseRussian              string = "пи-пи-пи"
	phraseSpanish              string = "piruleta"
	phraseSimplifiedChinese    string = "就说吧"
	phraseTraditionalChinese   string = "就說吧"
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
