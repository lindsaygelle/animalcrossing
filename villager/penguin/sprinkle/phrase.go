package sprinkle

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "frappe"
	phraseCanadianFrench       string = "mon glaçon"
	phraseDutch                string = "frappé"
	phraseFrench               string = "mon glaçon"
	phraseGerman               string = "bibib"
	phraseItalian              string = "unz unz"
	phraseJapanese             string = "ヒヤリ"
	phraseLatinAmericanSpanish string = "frapé"
	phraseKorean               string = "꽁꽁"
	phraseRussian              string = "пенка"
	phraseSpanish              string = "esmoquin"
	phraseSimplifiedChinese    string = "凉爽"
	phraseTraditionalChinese   string = "涼爽"
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
