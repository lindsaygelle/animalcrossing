package marcie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pouches"
	phraseCanadianFrench       string = "wallaby"
	phraseDutch                string = "buidels"
	phraseFrench               string = "wallaby"
	phraseGerman               string = "hüpf"
	phraseItalian              string = "me oui"
	phraseJapanese             string = "いいのよ"
	phraseLatinAmericanSpanish string = "nainoná"
	phraseKorean               string = "조아요"
	phraseRussian              string = "кармашек"
	phraseSpanish              string = "nainoná"
	phraseSimplifiedChinese    string = "好唷"
	phraseTraditionalChinese   string = "好唷"
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
