package marina

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "blurp"
	phraseCanadianFrench       string = "beurp"
	phraseDutch                string = "blorp"
	phraseFrench               string = "beurp"
	phraseGerman               string = "blubblub"
	phraseItalian              string = "blurp"
	phraseJapanese             string = "きゃ"
	phraseLatinAmericanSpanish string = "bliup"
	phraseKorean               string = "캬캬"
	phraseRussian              string = "хлюп"
	phraseSpanish              string = "dospatas"
	phraseSimplifiedChinese    string = "咔"
	phraseTraditionalChinese   string = "咔"
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
