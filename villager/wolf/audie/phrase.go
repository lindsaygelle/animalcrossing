package audie

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "foxtrot"
	phraseCanadianFrench       string = "trottine"
	phraseDutch                string = "foxtrot"
	phraseFrench               string = "trottine"
	phraseGerman               string = "foxtrott"
	phraseItalian              string = "foxtrot"
	phraseJapanese             string = "アハッ"
	phraseLatinAmericanSpanish string = "ulalilá"
	phraseKorean               string = "아하핫"
	phraseRussian              string = "фокстрот"
	phraseSpanish              string = "maravilla"
	phraseSimplifiedChinese    string = "呀哈"
	phraseTraditionalChinese   string = "呀哈"
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
