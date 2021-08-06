package zell

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pronk"
	phraseCanadianFrench       string = "wapiti"
	phraseDutch                string = "gazellig"
	phraseFrench               string = "wapiti"
	phraseGerman               string = "sproing"
	phraseItalian              string = "arriba"
	phraseJapanese             string = "たしかに"
	phraseLatinAmericanSpanish string = "bambín"
	phraseKorean               string = "슴사사"
	phraseRussian              string = "прыг-прыг"
	phraseSpanish              string = "bambín"
	phraseSimplifiedChinese    string = "鹿营"
	phraseTraditionalChinese   string = "鹿營"
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
