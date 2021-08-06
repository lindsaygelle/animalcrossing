package ellie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "li'l one"
	phraseCanadianFrench       string = "sacrelotte"
	phraseDutch                string = "dreumes"
	phraseFrench               string = "sacrelotte"
	phraseGerman               string = "polter"
	phraseItalian              string = "caromé"
	phraseJapanese             string = "ララン"
	phraseLatinAmericanSpanish string = "fi-fiú"
	phraseKorean               string = "트랄라"
	phraseRussian              string = "кроха"
	phraseSpanish              string = "trompi"
	phraseSimplifiedChinese    string = "啦啷"
	phraseTraditionalChinese   string = "啦啷"
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
