package dobie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ohmmm"
	phraseCanadianFrench       string = "loudonc"
	phraseDutch                string = "eh"
	phraseFrench               string = "peuchère"
	phraseGerman               string = "heuuul"
	phraseItalian              string = "gnams"
	phraseJapanese             string = "ゴホゴホ"
	phraseLatinAmericanSpanish string = "ooomm"
	phraseKorean               string = "음음"
	phraseRussian              string = "ом-м-м..."
	phraseSpanish              string = "ooomm"
	phraseSimplifiedChinese    string = "咳呵"
	phraseTraditionalChinese   string = "咳呵"
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
