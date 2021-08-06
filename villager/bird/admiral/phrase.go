package admiral

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "aye aye"
	phraseCanadianFrench       string = "cornichon"
	phraseDutch                string = "yep yep"
	phraseFrench               string = "cornichon"
	phraseGerman               string = "oh, ja"
	phraseItalian              string = "cipissi"
	phraseJapanese             string = "ってか"
	phraseLatinAmericanSpanish string = "piuc-piuc"
	phraseKorean               string = "오예"
	phraseRussian              string = "чиф-чиф"
	phraseSpanish              string = "avechucho"
	phraseSimplifiedChinese    string = "贯彻"
	phraseTraditionalChinese   string = "貫徹"
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
