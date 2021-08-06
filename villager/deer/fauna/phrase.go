package fauna

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "dearie"
	phraseCanadianFrench       string = "toudoux"
	phraseDutch                string = "do ree mi"
	phraseFrench               string = "toudoux"
	phraseGerman               string = "bimbam"
	phraseItalian              string = "vule vu"
	phraseJapanese             string = "でしか"
	phraseLatinAmericanSpanish string = "puchu"
	phraseKorean               string = "파샵파샵"
	phraseRussian              string = "олешек"
	phraseSpanish              string = "puchu"
	phraseSimplifiedChinese    string = "小鹿"
	phraseTraditionalChinese   string = "小鹿"
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
