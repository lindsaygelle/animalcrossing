package wade

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "so it goes"
	phraseCanadianFrench       string = "blanbec"
	phraseDutch                string = "zo is 't"
	phraseFrench               string = "blanbec"
	phraseGerman               string = "wackel"
	phraseItalian              string = "uhiii uuuh"
	phraseJapanese             string = "だからね"
	phraseLatinAmericanSpanish string = "espoink"
	phraseKorean               string = "뒤집어"
	phraseRussian              string = "вот так вот"
	phraseSpanish              string = "personaje"
	phraseSimplifiedChinese    string = "所以啦"
	phraseTraditionalChinese   string = "所以啦"
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
