package tiffany

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bunbun"
	phraseCanadianFrench       string = "pinpin"
	phraseDutch                string = "knijnelijn"
	phraseFrench               string = "pinpin"
	phraseGerman               string = "öhrchen"
	phraseItalian              string = "bunbun"
	phraseJapanese             string = "ってさ"
	phraseLatinAmericanSpanish string = "toin-lalá"
	phraseKorean               string = "라던데"
	phraseRussian              string = "зайчонок"
	phraseSpanish              string = "pompón"
	phraseSimplifiedChinese    string = "我说呀"
	phraseTraditionalChinese   string = "我說呀"
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
