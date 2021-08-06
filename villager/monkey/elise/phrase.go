package elise

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "puh-lease"
	phraseCanadianFrench       string = "steup"
	phraseDutch                string = "kom op zeg"
	phraseFrench               string = "steup"
	phraseGerman               string = "strebstreb"
	phraseItalian              string = "per favore"
	phraseJapanese             string = "だモン"
	phraseLatinAmericanSpanish string = "uh-lalá"
	phraseKorean               string = "몽몽"
	phraseRussian              string = "само собой"
	phraseSpanish              string = "ricura"
	phraseSimplifiedChinese    string = "孟孟"
	phraseTraditionalChinese   string = "孟孟"
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
