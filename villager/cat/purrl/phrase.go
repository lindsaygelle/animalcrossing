package purrl

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "kitten"
	phraseCanadianFrench       string = "chaton"
	phraseDutch                string = "kitten"
	phraseFrench               string = "chaton"
	phraseGerman               string = "kätzchen"
	phraseItalian              string = "micetto"
	phraseJapanese             string = "ふんっ"
	phraseLatinAmericanSpanish string = "michimiu"
	phraseKorean               string = "냠"
	phraseRussian              string = "кис-кис"
	phraseSpanish              string = "tontaina"
	phraseSimplifiedChinese    string = "嗯嗯"
	phraseTraditionalChinese   string = "嗯嗯"
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
