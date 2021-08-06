package rasher

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "swine"
	phraseCanadianFrench       string = "porcelet"
	phraseDutch                string = "big"
	phraseFrench               string = "porcelet"
	phraseGerman               string = "gronk"
	phraseItalian              string = "snoooink"
	phraseJapanese             string = "まんねん"
	phraseLatinAmericanSpanish string = "esnoink"
	phraseKorean               string = "꾸엑"
	phraseRussian              string = "хряк"
	phraseSpanish              string = "porcino"
	phraseSimplifiedChinese    string = "万年"
	phraseTraditionalChinese   string = "萬年"
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
