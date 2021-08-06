package annalisa

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "gumdrop"
	phraseCanadianFrench       string = "tap tap"
	phraseDutch                string = "dropje"
	phraseFrench               string = "tap tap"
	phraseGerman               string = "schnüffel"
	phraseItalian              string = "chic"
	phraseJapanese             string = "をかし"
	phraseLatinAmericanSpanish string = "snifffi"
	phraseKorean               string = "하양하양"
	phraseRussian              string = "жуй-жуй"
	phraseSpanish              string = "nubecillas"
	phraseSimplifiedChinese    string = "优雅"
	phraseTraditionalChinese   string = "優雅"
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
