package sheldon

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cardio"
	phraseCanadianFrench       string = "pacane"
	phraseDutch                string = "topfit"
	phraseFrench               string = "toutouffe"
	phraseGerman               string = "hoppi"
	phraseItalian              string = "flash"
	phraseJapanese             string = "クリクリ"
	phraseLatinAmericanSpanish string = "zasca"
	phraseKorean               string = "땡글땡글"
	phraseRussian              string = "треники"
	phraseSpanish              string = "zasca"
	phraseSimplifiedChinese    string = "鼓栗鼓栗"
	phraseTraditionalChinese   string = "鼓栗鼓栗"
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
