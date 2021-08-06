package rosie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "silly"
	phraseCanadianFrench       string = "flûte"
	phraseDutch                string = "gekkie"
	phraseFrench               string = "flûte"
	phraseGerman               string = "flöt"
	phraseItalian              string = "tontolon"
	phraseJapanese             string = "チェキ"
	phraseLatinAmericanSpanish string = "miaaau"
	phraseKorean               string = "헤이"
	phraseRussian              string = "рыбка"
	phraseSpanish              string = "miaaau"
	phraseSimplifiedChinese    string = "看看"
	phraseTraditionalChinese   string = "看看"
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
