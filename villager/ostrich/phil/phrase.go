package phil

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hurk"
	phraseCanadianFrench       string = "bécot"
	phraseDutch                string = "struis"
	phraseFrench               string = "bécot"
	phraseGerman               string = "zwoing"
	phraseItalian              string = "tambien"
	phraseJapanese             string = "ホロロ"
	phraseLatinAmericanSpanish string = "jurk"
	phraseKorean               string = "호롤로"
	phraseRussian              string = "хех"
	phraseSpanish              string = "jurk"
	phraseSimplifiedChinese    string = "轰隆隆"
	phraseTraditionalChinese   string = "轟隆隆"
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
