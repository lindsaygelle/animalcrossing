package sprocket

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "zort"
	phraseCanadianFrench       string = "pik pik"
	phraseDutch                string = "klik-kluk"
	phraseFrench               string = "pik pik"
	phraseGerman               string = "affenzahn"
	phraseItalian              string = "niiun"
	phraseJapanese             string = "だメカ"
	phraseLatinAmericanSpanish string = "chiuuun"
	phraseKorean               string = "치지직"
	phraseRussian              string = "дзынь"
	phraseSpanish              string = "chiuuun"
	phraseSimplifiedChinese    string = "机械"
	phraseTraditionalChinese   string = "機械"
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
