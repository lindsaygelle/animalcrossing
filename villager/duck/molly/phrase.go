package molly

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "quackidee"
	phraseCanadianFrench       string = "canénette"
	phraseDutch                string = "topper"
	phraseFrench               string = "canénette"
	phraseGerman               string = "krümel"
	phraseItalian              string = "coin coin"
	phraseJapanese             string = "そうかも"
	phraseLatinAmericanSpanish string = "cuáquidi"
	phraseKorean               string = "그럴지도"
	phraseRussian              string = "крякульки"
	phraseSpanish              string = "deverdá"
	phraseSimplifiedChinese    string = "说不定鸭"
	phraseTraditionalChinese   string = "說不定鴨"
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
