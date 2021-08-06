package chester

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "rookie"
	phraseCanadianFrench       string = "pla-pla"
	phraseDutch                string = "bamboe"
	phraseFrench               string = "pla-pla"
	phraseGerman               string = "baaambus"
	phraseItalian              string = "brupp"
	phraseJapanese             string = "だバンブ"
	phraseLatinAmericanSpanish string = "ahivá"
	phraseKorean               string = "부끄"
	phraseRussian              string = "бамбук"
	phraseSpanish              string = "ahivá"
	phraseSimplifiedChinese    string = "竹子"
	phraseTraditionalChinese   string = "竹子"
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
