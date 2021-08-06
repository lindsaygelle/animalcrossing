package agnes

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snuffle"
	phraseCanadianFrench       string = "sauciflard"
	phraseDutch                string = "snuffelaar"
	phraseFrench               string = "sauciflard"
	phraseGerman               string = "schnorrz"
	phraseItalian              string = "splosh"
	phraseJapanese             string = "ブフフ"
	phraseLatinAmericanSpanish string = "dindoink"
	phraseKorean               string = "꿀꾸루"
	phraseRussian              string = "хрю-хрю"
	phraseSpanish              string = "dindoink"
	phraseSimplifiedChinese    string = "噗呼呼"
	phraseTraditionalChinese   string = "噗呼呼"
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
