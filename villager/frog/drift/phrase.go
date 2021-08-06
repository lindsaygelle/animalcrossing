package drift

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "dribbit"
	phraseCanadianFrench       string = "braaaa"
	phraseDutch                string = "gozer"
	phraseFrench               string = "braaaa"
	phraseGerman               string = "quak"
	phraseItalian              string = "girin"
	phraseJapanese             string = "ゲロゲロ"
	phraseLatinAmericanSpanish string = "dribit"
	phraseKorean               string = "개굴개굴"
	phraseRussian              string = "бро"
	phraseSpanish              string = "colega"
	phraseSimplifiedChinese    string = "聒聒"
	phraseTraditionalChinese   string = "嘓嘓"
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
