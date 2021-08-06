package rudy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mush"
	phraseCanadianFrench       string = "mouah"
	phraseDutch                string = "robijn"
	phraseFrench               string = "alatienne"
	phraseGerman               string = "katzaaa"
	phraseItalian              string = "ron ron"
	phraseJapanese             string = "とかナ"
	phraseLatinAmericanSpanish string = "cachusco"
	phraseKorean               string = "그러거나"
	phraseRussian              string = "шмяк"
	phraseSpanish              string = "cachusco"
	phraseSimplifiedChinese    string = "之类的"
	phraseTraditionalChinese   string = "之類的"
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
