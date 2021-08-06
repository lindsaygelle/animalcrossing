package twiggy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cheepers"
	phraseCanadianFrench       string = "coolos"
	phraseDutch                string = "fwieeet"
	phraseFrench               string = "coolos"
	phraseGerman               string = "zwirtschi"
	phraseItalian              string = "ciiip"
	phraseJapanese             string = "ッピ"
	phraseLatinAmericanSpanish string = "tirití"
	phraseKorean               string = "크"
	phraseRussian              string = "тю-вить"
	phraseSpanish              string = "tirití"
	phraseSimplifiedChinese    string = "叽"
	phraseTraditionalChinese   string = "嘰"
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
