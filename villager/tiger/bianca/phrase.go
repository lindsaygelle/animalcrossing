package bianca

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "glimmer"
	phraseCanadianFrench       string = "la vie"
	phraseDutch                string = "witje"
	phraseFrench               string = "croky"
	phraseGerman               string = "prrr"
	phraseItalian              string = "yeppa"
	phraseJapanese             string = "でヒョウ"
	phraseLatinAmericanSpanish string = "ruarri"
	phraseKorean               string = "그래호"
	phraseRussian              string = "хлоп-хлоп"
	phraseSpanish              string = "guiñito"
	phraseSimplifiedChinese    string = "雪豹"
	phraseTraditionalChinese   string = "雪豹"
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
