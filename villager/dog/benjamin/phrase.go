package benjamin

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "alrighty"
	phraseCanadianFrench       string = "croquettes"
	phraseDutch                string = "oké dan"
	phraseFrench               string = "croquettes"
	phraseGerman               string = "winsel"
	phraseItalian              string = "bauuau"
	phraseJapanese             string = "ではでは"
	phraseLatinAmericanSpanish string = "tirirí"
	phraseKorean               string = "그럼그럼"
	phraseRussian              string = "пор-рядок"
	phraseSpanish              string = "tirirí"
	phraseSimplifiedChinese    string = "那么那么"
	phraseTraditionalChinese   string = "那麼那麼"
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
