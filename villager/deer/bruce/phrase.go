package bruce

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "gruff"
	phraseCanadianFrench       string = "ouain"
	phraseDutch                string = "burl"
	phraseFrench               string = "nunul"
	phraseGerman               string = "röööhr"
	phraseItalian              string = "uauwaii"
	phraseJapanese             string = "しかしな"
	phraseLatinAmericanSpanish string = "aquesí"
	phraseKorean               string = "아비옹"
	phraseRussian              string = "забияка"
	phraseSpanish              string = "aquesí"
	phraseSimplifiedChinese    string = "福鹿寿"
	phraseTraditionalChinese   string = "福鹿壽"
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
