package rodney

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "le ham"
	phraseCanadianFrench       string = "yes sir"
	phraseDutch                string = "le ham"
	phraseFrench               string = "caries"
	phraseGerman               string = "mjamjam"
	phraseItalian              string = "flop"
	phraseJapanese             string = "ヒュー"
	phraseLatinAmericanSpanish string = "adoya"
	phraseKorean               string = "휴우"
	phraseRussian              string = "ле хом"
	phraseSpanish              string = "adoya"
	phraseSimplifiedChinese    string = "咻咻"
	phraseTraditionalChinese   string = "咻咻"
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
