package prince

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "burrup"
	phraseCanadianFrench       string = "boiiiiing"
	phraseDutch                string = "blurp"
	phraseFrench               string = "boiiiiing"
	phraseGerman               string = "gurp"
	phraseItalian              string = "berp"
	phraseJapanese             string = "ですだ"
	phraseLatinAmericanSpanish string = "buuurp"
	phraseKorean               string = "이리오슈"
	phraseRussian              string = "квак-квак"
	phraseSpanish              string = "buuurp"
	phraseSimplifiedChinese    string = "是的"
	phraseTraditionalChinese   string = "是的"
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
