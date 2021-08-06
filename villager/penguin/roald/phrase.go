package roald

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "b-b-buddy"
	phraseCanadianFrench       string = "pépépère"
	phraseDutch                string = "m-m-maatje"
	phraseFrench               string = "pépépère"
	phraseGerman               string = "k-k-kumpel"
	phraseItalian              string = "br-br-brr"
	phraseJapanese             string = "だペン"
	phraseLatinAmericanSpanish string = "titirito"
	phraseKorean               string = "팽팽"
	phraseRussian              string = "д-дружище"
	phraseSpanish              string = "titirito"
	phraseSimplifiedChinese    string = "企鹅"
	phraseTraditionalChinese   string = "企鵝"
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
