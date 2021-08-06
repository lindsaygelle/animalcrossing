package tabby

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "me-WOW"
	phraseCanadianFrench       string = "mi-AOUUH"
	phraseDutch                string = "mi-WAUW"
	phraseFrench               string = "mi-AOUUH"
	phraseGerman               string = "mrriau"
	phraseItalian              string = "meOW"
	phraseJapanese             string = "にゃは"
	phraseLatinAmericanSpanish string = "animiaul"
	phraseKorean               string = "삐뽀"
	phraseRussian              string = "мя-ого"
	phraseSpanish              string = "animiaul"
	phraseSimplifiedChinese    string = "喵吼"
	phraseTraditionalChinese   string = "喵吼"
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
