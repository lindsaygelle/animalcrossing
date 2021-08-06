package bunnie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "tee-hee"
	phraseCanadianFrench       string = "hé hé hé"
	phraseDutch                string = "hihi"
	phraseFrench               string = "hé hé hé"
	phraseGerman               string = "tiihii"
	phraseItalian              string = "tiptap"
	phraseJapanese             string = "みたいな"
	phraseLatinAmericanSpanish string = "oy-oy-oy"
	phraseKorean               string = "그렇지뭐"
	phraseRussian              string = "хи-хи"
	phraseSpanish              string = "oy-oy-oy"
	phraseSimplifiedChinese    string = "好像喔"
	phraseTraditionalChinese   string = "好像喔"
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
