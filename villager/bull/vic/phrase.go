package vic

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cud"
	phraseCanadianFrench       string = "moumoul"
	phraseDutch                string = "kauw"
	phraseFrench               string = "moumoul"
	phraseGerman               string = "grumpf"
	phraseItalian              string = "trullalà"
	phraseJapanese             string = "モイ"
	phraseLatinAmericanSpanish string = "brufff"
	phraseKorean               string = "뭐임"
	phraseRussian              string = "полундра"
	phraseSpanish              string = "muuubién"
	phraseSimplifiedChinese    string = "哞"
	phraseTraditionalChinese   string = "哞"
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
