package curlos

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "shearly"
	phraseCanadianFrench       string = "toison"
	phraseDutch                string = "scheerlijk"
	phraseFrench               string = "toison"
	phraseGerman               string = "zottel"
	phraseItalian              string = "hermoso"
	phraseJapanese             string = "ベイビー"
	phraseLatinAmericanSpanish string = "laaaná"
	phraseKorean               string = "자기야"
	phraseRussian              string = "завитушка"
	phraseSpanish              string = "lanitas"
	phraseSimplifiedChinese    string = "北鼻"
	phraseTraditionalChinese   string = "北鼻"
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
