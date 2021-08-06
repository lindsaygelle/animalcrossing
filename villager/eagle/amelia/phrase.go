package amelia

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "cuz"
	phraseCanadianFrench       string = "croooa"
	phraseDutch                string = "broeder"
	phraseFrench               string = "croooa"
	phraseGerman               string = "adlerauge"
	phraseItalian              string = "mini mini"
	phraseJapanese             string = "カラカラ"
	phraseLatinAmericanSpanish string = "cro-ah"
	phraseKorean               string = "엄멈머"
	phraseRussian              string = "вот"
	phraseSpanish              string = "chinchilla"
	phraseSimplifiedChinese    string = "卡拉卡拉"
	phraseTraditionalChinese   string = "卡拉卡拉"
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
