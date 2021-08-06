package eunice

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "lambchop"
	phraseCanadianFrench       string = "bêbêêêêh"
	phraseDutch                string = "krulletje"
	phraseFrench               string = "bêbêêêêh"
	phraseGerman               string = "la-a-amm"
	phraseItalian              string = "laaagna"
	phraseJapanese             string = "メェー"
	phraseLatinAmericanSpanish string = "yupitiii"
	phraseKorean               string = "메에에"
	phraseRussian              string = "ребрышко"
	phraseSpanish              string = "veeenga"
	phraseSimplifiedChinese    string = "咩"
	phraseTraditionalChinese   string = "咩"
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
