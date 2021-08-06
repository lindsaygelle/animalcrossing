package flo

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cha"
	phraseCanadianFrench       string = "frigogo"
	phraseDutch                string = "duh"
	phraseFrench               string = "golri"
	phraseGerman               string = "flatsch"
	phraseItalian              string = "iglù"
	phraseJapanese             string = "じゃんよ"
	phraseLatinAmericanSpanish string = "fresqui"
	phraseKorean               string = "맞잖아"
	phraseRussian              string = "чхи"
	phraseSpanish              string = "fresqui"
	phraseSimplifiedChinese    string = "这样唷"
	phraseTraditionalChinese   string = "這樣唷"
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
