package apple

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cheekers"
	phraseCanadianFrench       string = "bibille"
	phraseDutch                string = "wangetjes"
	phraseFrench               string = "bibille"
	phraseGerman               string = "fiep"
	phraseItalian              string = "triplo uau"
	phraseJapanese             string = "キュルン"
	phraseLatinAmericanSpanish string = "do-re-mi"
	phraseKorean               string = "큐룽"
	phraseRussian              string = "щечки"
	phraseSpanish              string = "do-re-mi"
	phraseSimplifiedChinese    string = "转转"
	phraseTraditionalChinese   string = "轉轉"
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
