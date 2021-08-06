package monique

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pffffft"
	phraseCanadianFrench       string = "pffffft"
	phraseDutch                string = "pffffft"
	phraseFrench               string = "pffffft"
	phraseGerman               string = "mrrrrrrr"
	phraseItalian              string = "miaù"
	phraseJapanese             string = "ウフーン"
	phraseLatinAmericanSpanish string = "ronrón"
	phraseKorean               string = "우"
	phraseRussian              string = "фи-фи"
	phraseSpanish              string = "ronrón"
	phraseSimplifiedChinese    string = "嗯哼"
	phraseTraditionalChinese   string = "嗯哼"
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
