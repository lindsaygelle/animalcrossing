package plucky

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chicky-poo"
	phraseCanadianFrench       string = "galine"
	phraseDutch                string = "kippetje"
	phraseFrench               string = "galine"
	phraseGerman               string = "hühnermist"
	phraseItalian              string = "zampoli"
	phraseJapanese             string = "どうだい"
	phraseLatinAmericanSpanish string = "corocó"
	phraseKorean               string = "그게어때"
	phraseRussian              string = "щип-щип"
	phraseSpanish              string = "tesorito"
	phraseSimplifiedChinese    string = "泰泰"
	phraseTraditionalChinese   string = "泰泰"
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
