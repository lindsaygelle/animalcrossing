package cesar

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "highness"
	phraseCanadianFrench       string = "brutus"
	phraseDutch                string = "hoogheid"
	phraseFrench               string = "brutus"
	phraseGerman               string = "hoheit"
	phraseItalian              string = "dado"
	phraseJapanese             string = "ウホウホ"
	phraseLatinAmericanSpanish string = "ejeeem"
	phraseKorean               string = "우갸우갸"
	phraseRussian              string = "высочество"
	phraseSpanish              string = "monigote"
	phraseSimplifiedChinese    string = "拍胸脯"
	phraseTraditionalChinese   string = "拍胸脯"
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
