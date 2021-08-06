package kitt

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "child"
	phraseCanadianFrench       string = "gamin"
	phraseDutch                string = "kind"
	phraseFrench               string = "gamin"
	phraseGerman               string = "kindchen"
	phraseItalian              string = "cucciolo"
	phraseJapanese             string = "ッポン"
	phraseLatinAmericanSpanish string = "chiquichí"
	phraseKorean               string = "폴짝"
	phraseRussian              string = "детка"
	phraseSpanish              string = "mi alma"
	phraseSimplifiedChinese    string = "补补"
	phraseTraditionalChinese   string = "補補"
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
