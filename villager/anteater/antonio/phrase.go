package antonio

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "honk"
	phraseCanadianFrench       string = "pouet"
	phraseDutch                string = "snuit"
	phraseFrench               string = "pouet"
	phraseGerman               string = "schlürrrf"
	phraseItalian              string = "honk"
	phraseJapanese             string = "ホントに"
	phraseLatinAmericanSpanish string = "fufuf"
	phraseKorean               string = "진짜로"
	phraseRussian              string = "го-го-го"
	phraseSpanish              string = "fufuf"
	phraseSimplifiedChinese    string = "真诚"
	phraseTraditionalChinese   string = "真誠"
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
