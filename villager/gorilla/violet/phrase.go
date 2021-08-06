package violet

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sweetie"
	phraseCanadianFrench       string = "gong"
	phraseDutch                string = "lieverd"
	phraseFrench               string = "gong"
	phraseGerman               string = "auauauu"
	phraseItalian              string = "tump tump"
	phraseJapanese             string = "アイヤ"
	phraseLatinAmericanSpanish string = "uh-uh"
	phraseKorean               string = "에그머닛"
	phraseRussian              string = "сластена"
	phraseSpanish              string = "canapé"
	phraseSimplifiedChinese    string = "唉呀"
	phraseTraditionalChinese   string = "唉呀"
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
