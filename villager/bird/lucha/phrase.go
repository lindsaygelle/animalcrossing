package lucha

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cacaw"
	phraseCanadianFrench       string = "yaaaah"
	phraseDutch                string = "kakauw"
	phraseFrench               string = "yaaaah"
	phraseGerman               string = "mauser"
	phraseItalian              string = "lorito"
	phraseJapanese             string = "カァー"
	phraseLatinAmericanSpanish string = "pío"
	phraseKorean               string = "까아악"
	phraseRussian              string = "кар-ко"
	phraseSpanish              string = "picopico"
	phraseSimplifiedChinese    string = "嘎"
	phraseTraditionalChinese   string = "嘎"
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
