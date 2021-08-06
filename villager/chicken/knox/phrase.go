package knox

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cluckling"
	phraseCanadianFrench       string = "ra-ta-tak"
	phraseDutch                string = "nestei"
	phraseFrench               string = "ra-ta-tak"
	phraseGerman               string = "knusprig"
	phraseItalian              string = "uovole"
	phraseJapanese             string = "せいッ"
	phraseLatinAmericanSpanish string = "cocorocó"
	phraseKorean               string = "하앗"
	phraseRussian              string = "короко"
	phraseSpanish              string = "mazorca"
	phraseSimplifiedChinese    string = "等待"
	phraseTraditionalChinese   string = "等待"
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
