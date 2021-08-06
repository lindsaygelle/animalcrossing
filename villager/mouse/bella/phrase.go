package bella

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "eeks"
	phraseCanadianFrench       string = "hiiiii"
	phraseDutch                string = "ieps"
	phraseFrench               string = "hiiiii"
	phraseGerman               string = "visavisa"
	phraseItalian              string = "eeks"
	phraseJapanese             string = "ギャハッ"
	phraseLatinAmericanSpanish string = "petisú"
	phraseKorean               string = "캬학"
	phraseRussian              string = "пип"
	phraseSpanish              string = "petisú"
	phraseSimplifiedChinese    string = "嘎哈"
	phraseTraditionalChinese   string = "嘎哈"
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
