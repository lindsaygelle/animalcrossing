package greta

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "yelp"
	phraseCanadianFrench       string = "okka"
	phraseDutch                string = "piepzak"
	phraseFrench               string = "téma"
	phraseGerman               string = "nagnag"
	phraseItalian              string = "rattaplan"
	phraseJapanese             string = "おほほ"
	phraseLatinAmericanSpanish string = "ñiñi-ñiá"
	phraseKorean               string = "오호호"
	phraseRussian              string = "писк"
	phraseSpanish              string = "bocadito"
	phraseSimplifiedChinese    string = "呵呵呵"
	phraseTraditionalChinese   string = "呵呵呵"
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
