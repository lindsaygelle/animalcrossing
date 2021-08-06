package cheri

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tralala"
	phraseCanadianFrench       string = "tralala"
	phraseDutch                string = "tralala"
	phraseFrench               string = "tralala"
	phraseGerman               string = "tralala"
	phraseItalian              string = "trallalà"
	phraseJapanese             string = "なんてね"
	phraseLatinAmericanSpanish string = "tralará"
	phraseKorean               string = "어쩜이래"
	phraseRussian              string = "тра-ля-ля"
	phraseSpanish              string = "tralará"
	phraseSimplifiedChinese    string = "开玩笑的"
	phraseTraditionalChinese   string = "開玩笑的"
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
