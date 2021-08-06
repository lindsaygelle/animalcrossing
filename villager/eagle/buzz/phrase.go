package buzz

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "captain"
	phraseCanadianFrench       string = "flap flap"
	phraseDutch                string = "kaptein"
	phraseFrench               string = "flap flap"
	phraseGerman               string = "käpten"
	phraseItalian              string = "capo"
	phraseJapanese             string = "ッキーッ"
	phraseLatinAmericanSpanish string = "capííí"
	phraseKorean               string = "짜란～"
	phraseRussian              string = "капитан"
	phraseSpanish              string = "cadete"
	phraseSimplifiedChinese    string = "继继"
	phraseTraditionalChinese   string = "繼繼"
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
