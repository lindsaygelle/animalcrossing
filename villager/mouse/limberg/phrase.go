package limberg

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "squinky"
	phraseCanadianFrench       string = "scriiich"
	phraseDutch                string = "piepsels"
	phraseFrench               string = "scriiich"
	phraseGerman               string = "pfiffikus"
	phraseItalian              string = "squick"
	phraseJapanese             string = "てやんで"
	phraseLatinAmericanSpanish string = "scroch"
	phraseKorean               string = "긍가벼"
	phraseRussian              string = "фью-фью"
	phraseSpanish              string = "caramba"
	phraseSimplifiedChinese    string = "非也"
	phraseTraditionalChinese   string = "非也"
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
