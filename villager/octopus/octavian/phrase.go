package octavian

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sucker"
	phraseCanadianFrench       string = "chenapan"
	phraseDutch                string = "zuignap"
	phraseFrench               string = "chenapan"
	phraseGerman               string = "schling"
	phraseItalian              string = "perdindi"
	phraseJapanese             string = "タコ"
	phraseLatinAmericanSpanish string = "glop-glop"
	phraseKorean               string = "쭉쭉"
	phraseRussian              string = "все пропало"
	phraseSpanish              string = "chupatinta"
	phraseSimplifiedChinese    string = "章鱼"
	phraseTraditionalChinese   string = "章魚"
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
