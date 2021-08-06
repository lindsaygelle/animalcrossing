package cube

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "brainfreeze"
	phraseCanadianFrench       string = "cococopain"
	phraseDutch                string = "k-k-kerel"
	phraseFrench               string = "cococopain"
	phraseGerman               string = "d-d-dude"
	phraseItalian              string = "brivido"
	phraseJapanese             string = "ペンペン"
	phraseLatinAmericanSpanish string = "waap"
	phraseKorean               string = "땡땡"
	phraseRussian              string = "ч-чубрик"
	phraseSpanish              string = "waap"
	phraseSimplifiedChinese    string = "十字"
	phraseTraditionalChinese   string = "十字"
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
