package zucker

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bloop"
	phraseCanadianFrench       string = "sploush"
	phraseDutch                string = "bloep"
	phraseFrench               string = "sploush"
	phraseGerman               string = "saug"
	phraseItalian              string = "blub"
	phraseJapanese             string = "せやねん"
	phraseLatinAmericanSpanish string = "churubú"
	phraseKorean               string = "약히"
	phraseRussian              string = "буль"
	phraseSpanish              string = "churubú"
	phraseSimplifiedChinese    string = "认同"
	phraseTraditionalChinese   string = "認同"
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
