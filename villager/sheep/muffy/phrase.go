package muffy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "nightshade"
	phraseCanadianFrench       string = "tricoti"
	phraseDutch                string = "nachtvacht"
	phraseFrench               string = "tricoti"
	phraseGerman               string = "knuffi"
	phraseItalian              string = "lanadà"
	phraseJapanese             string = "おつかれ"
	phraseLatinAmericanSpanish string = "lanavá"
	phraseKorean               string = "수고"
	phraseRussian              string = "травинка"
	phraseSpanish              string = "lanas"
	phraseSimplifiedChinese    string = "辛苦了"
	phraseTraditionalChinese   string = "辛苦了"
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
