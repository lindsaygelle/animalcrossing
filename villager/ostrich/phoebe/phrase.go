package phoebe

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sparky"
	phraseCanadianFrench       string = "chanmé"
	phraseDutch                string = "vonkje"
	phraseFrench               string = "chanmé"
	phraseGerman               string = "pieks"
	phraseItalian              string = "behmù"
	phraseJapanese             string = "アツイわ"
	phraseLatinAmericanSpanish string = "equilicuá"
	phraseKorean               string = "아뜨뜨"
	phraseRussian              string = "искорка"
	phraseSpanish              string = "equilicuá"
	phraseSimplifiedChinese    string = "好热"
	phraseTraditionalChinese   string = "好熱"
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
