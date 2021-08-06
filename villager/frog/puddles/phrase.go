package puddles

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "splish"
	phraseCanadianFrench       string = "splufff"
	phraseDutch                string = "plons"
	phraseFrench               string = "splufff"
	phraseGerman               string = "platschi"
	phraseItalian              string = "splash"
	phraseJapanese             string = "っちゃ"
	phraseLatinAmericanSpanish string = "esplís"
	phraseKorean               string = "그랬쪄"
	phraseRussian              string = "плюх"
	phraseSpanish              string = "esplís"
	phraseSimplifiedChinese    string = "石头"
	phraseTraditionalChinese   string = "石頭"
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
