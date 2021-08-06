package spike

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "punk"
	phraseCanadianFrench       string = "groumph"
	phraseDutch                string = "schelm"
	phraseFrench               string = "saperlotte"
	phraseGerman               string = "punk"
	phraseItalian              string = "pachi"
	phraseJapanese             string = "でヤンキ"
	phraseLatinAmericanSpanish string = "papúm"
	phraseKorean               string = "빠직"
	phraseRussian              string = "бродяга"
	phraseSpanish              string = "granuja"
	phraseSimplifiedChinese    string = "混混"
	phraseTraditionalChinese   string = "混混"
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
