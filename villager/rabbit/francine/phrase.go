package francine

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "karat"
	phraseCanadianFrench       string = "carotte"
	phraseDutch                string = "wortels"
	phraseFrench               string = "carotte"
	phraseGerman               string = "hey man"
	phraseItalian              string = "lulalà"
	phraseJapanese             string = "ルララ"
	phraseLatinAmericanSpanish string = "kaninchen"
	phraseKorean               string = "쇼봉"
	phraseRussian              string = "морковка"
	phraseSpanish              string = "kaninchen"
	phraseSimplifiedChinese    string = "噜啦啦"
	phraseTraditionalChinese   string = "嚕啦啦"
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
