package rod

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ace"
	phraseCanadianFrench       string = "tip top"
	phraseDutch                string = "toppie"
	phraseFrench               string = "tip top"
	phraseGerman               string = "boa-ey"
	phraseItalian              string = "flapflap"
	phraseJapanese             string = "すっげぇ"
	phraseLatinAmericanSpanish string = "chopchop"
	phraseKorean               string = "무진장"
	phraseRussian              string = "мастерски"
	phraseSpanish              string = "quesito"
	phraseSimplifiedChinese    string = "厉害"
	phraseTraditionalChinese   string = "厲害"
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
