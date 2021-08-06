package aurora

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "b-b-baby"
	phraseCanadianFrench       string = "b-b-bébé"
	phraseDutch                string = "b-b-baby"
	phraseFrench               string = "b-b-bébé"
	phraseGerman               string = "b-bestens"
	phraseItalian              string = "b-b-baby"
	phraseJapanese             string = "だジョー"
	phraseLatinAmericanSpanish string = "pescadí"
	phraseKorean               string = "랄라"
	phraseRussian              string = "к-к-крошка"
	phraseSpanish              string = "pescadito"
	phraseSimplifiedChinese    string = "若"
	phraseTraditionalChinese   string = "若"
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
