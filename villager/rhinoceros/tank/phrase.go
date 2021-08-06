package tank

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "kerPOW"
	phraseCanadianFrench       string = "kaboum"
	phraseDutch                string = "klaBAM"
	phraseFrench               string = "kaboum"
	phraseGerman               string = "rumpel"
	phraseItalian              string = "zamperù"
	phraseJapanese             string = "ですサイ"
	phraseLatinAmericanSpanish string = "tapumba"
	phraseKorean               string = "아뿔소"
	phraseRussian              string = "бабах"
	phraseSpanish              string = "tapumba"
	phraseSimplifiedChinese    string = "犀犀"
	phraseTraditionalChinese   string = "犀犀"
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
