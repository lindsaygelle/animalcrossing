package pierce

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hawkeye"
	phraseCanadianFrench       string = "tout vu"
	phraseDutch                string = "arendsoog"
	phraseFrench               string = "œil agile"
	phraseGerman               string = "schwinge"
	phraseItalian              string = "pronti via"
	phraseJapanese             string = "バサバサ"
	phraseLatinAmericanSpanish string = "aguilííí"
	phraseKorean               string = "퍼덕퍼덕"
	phraseRussian              string = "глаз-алмаз"
	phraseSpanish              string = "aguililla"
	phraseSimplifiedChinese    string = "啪沙啪沙"
	phraseTraditionalChinese   string = "啪沙啪沙"
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
