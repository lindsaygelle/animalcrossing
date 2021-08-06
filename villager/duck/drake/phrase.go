package drake

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "quacko"
	phraseCanadianFrench       string = "ben couac"
	phraseDutch                string = "snabbel"
	phraseFrench               string = "canard"
	phraseGerman               string = "schnabbl"
	phraseItalian              string = "quaqua"
	phraseJapanese             string = "かもね"
	phraseLatinAmericanSpanish string = "cui-cuac"
	phraseKorean               string = "어쩌면"
	phraseRussian              string = "кра-а"
	phraseSpanish              string = "polluelo"
	phraseSimplifiedChinese    string = "或许鸭"
	phraseTraditionalChinese   string = "或許鴨"
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
