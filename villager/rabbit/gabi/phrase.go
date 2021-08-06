package gabi

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "honeybun"
	phraseCanadianFrench       string = "mon lièvre"
	phraseDutch                string = "honnepon"
	phraseFrench               string = "mon lièvre"
	phraseGerman               string = "liebling"
	phraseItalian              string = "mielito"
	phraseJapanese             string = "だっち"
	phraseLatinAmericanSpanish string = "reboing"
	phraseKorean               string = "꺄오"
	phraseRussian              string = "зайчик мой"
	phraseSpanish              string = "mari"
	phraseSimplifiedChinese    string = "琪他"
	phraseTraditionalChinese   string = "琪他"
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
