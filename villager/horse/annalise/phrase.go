package annalise

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nipper"
	phraseCanadianFrench       string = "hippique"
	phraseDutch                string = "veulen"
	phraseFrench               string = "galapia"
	phraseGerman               string = "mausezahn"
	phraseItalian              string = "palio"
	phraseJapanese             string = "サヴァ"
	phraseLatinAmericanSpanish string = "nih-nah"
	phraseKorean               string = "봉쥬르"
	phraseRussian              string = "кусь-кусь"
	phraseSpanish              string = "criatura"
	phraseSimplifiedChinese    string = "舒服"
	phraseTraditionalChinese   string = "舒服"
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
