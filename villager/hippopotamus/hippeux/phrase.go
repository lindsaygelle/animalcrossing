package hippeux

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "natch"
	phraseCanadianFrench       string = "tam-tam"
	phraseDutch                string = "tuurlijk"
	phraseFrench               string = "tam-tam"
	phraseGerman               string = "hipphipp"
	phraseItalian              string = "vale"
	phraseJapanese             string = "イェス"
	phraseLatinAmericanSpanish string = "chofchof"
	phraseKorean               string = "예스"
	phraseRussian              string = "естессно"
	phraseSpanish              string = "chofchof"
	phraseSimplifiedChinese    string = "Yes"
	phraseTraditionalChinese   string = "Yes"
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
