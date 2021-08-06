package jeremiah

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nee-deep"
	phraseCanadianFrench       string = "brooap"
	phraseDutch                string = "rikkekik"
	phraseFrench               string = "brooap"
	phraseGerman               string = "nidip"
	phraseItalian              string = "fri fri"
	phraseJapanese             string = "にゃむ"
	phraseLatinAmericanSpanish string = "croc-croc"
	phraseKorean               string = "딩동댕"
	phraseRussian              string = "буль-буль"
	phraseSpanish              string = "croc-croc"
	phraseSimplifiedChinese    string = "嚼嚼"
	phraseTraditionalChinese   string = "嚼嚼"
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
