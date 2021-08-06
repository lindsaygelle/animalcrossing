package flora

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pinky"
	phraseCanadianFrench       string = "rosi rosa"
	phraseDutch                string = "roze"
	phraseFrench               string = "guiboles"
	phraseGerman               string = "flappflapp"
	phraseItalian              string = "sivuplé"
	phraseJapanese             string = "だミン"
	phraseLatinAmericanSpanish string = "croqui"
	phraseKorean               string = "밍고밍고"
	phraseRussian              string = "розочка"
	phraseSpanish              string = "croqui"
	phraseSimplifiedChinese    string = "红鹤"
	phraseTraditionalChinese   string = "紅鶴"
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
