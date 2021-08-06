package tangy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "reeeeOWR"
	phraseCanadianFrench       string = "roooaR"
	phraseDutch                string = "sinas"
	phraseFrench               string = "roooaR"
	phraseGerman               string = "grraul"
	phraseItalian              string = "miiiiaOU"
	phraseJapanese             string = "みかん"
	phraseLatinAmericanSpanish string = "mirriau"
	phraseKorean               string = "귤귤"
	phraseRussian              string = "фруктик"
	phraseSpanish              string = "mandarinas"
	phraseSimplifiedChinese    string = "蜜柑"
	phraseTraditionalChinese   string = "蜜柑"
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
