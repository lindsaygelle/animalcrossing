package derwin

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "derrrr"
	phraseCanadianFrench       string = "derrrr"
	phraseDutch                string = "ergo"
	phraseFrench               string = "derrrr"
	phraseGerman               string = "krook"
	phraseItalian              string = "quaccolo"
	phraseJapanese             string = "ね！ママ"
	phraseLatinAmericanSpanish string = "cuato"
	phraseKorean               string = "엄마"
	phraseRussian              string = "кррря"
	phraseSpanish              string = "cuato"
	phraseSimplifiedChinese    string = "喂！妈妈"
	phraseTraditionalChinese   string = "喂！媽媽"
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
