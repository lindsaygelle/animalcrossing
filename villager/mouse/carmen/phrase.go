package carmen

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nougat"
	phraseCanadianFrench       string = "nougat"
	phraseDutch                string = "noga"
	phraseFrench               string = "pinou-nou"
	phraseGerman               string = "langohr"
	phraseItalian              string = "gnuf gnuf"
	phraseJapanese             string = "まじで"
	phraseLatinAmericanSpanish string = "nufinuf"
	phraseKorean               string = "진짜진짜"
	phraseRussian              string = "карамелька"
	phraseSpanish              string = "arcoíris"
	phraseSimplifiedChinese    string = "真的"
	phraseTraditionalChinese   string = "真的"
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
