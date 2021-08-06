package lucy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "snoooink"
	phraseCanadianFrench       string = "groingroin"
	phraseDutch                string = "knorrie"
	phraseFrench               string = "groingroin"
	phraseGerman               string = "schnoink"
	phraseItalian              string = "snoink"
	phraseJapanese             string = "よぅ"
	phraseLatinAmericanSpanish string = "cochín"
	phraseKorean               string = "예예"
	phraseRussian              string = "хрю-ю-ю"
	phraseSpanish              string = "cochín"
	phraseSimplifiedChinese    string = "唷"
	phraseTraditionalChinese   string = "唷"
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
