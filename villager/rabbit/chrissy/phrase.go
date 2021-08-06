package chrissy

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sparkles"
	phraseCanadianFrench       string = "hop"
	phraseDutch                string = "sprankel"
	phraseFrench               string = "hop"
	phraseGerman               string = "knuddl"
	phraseItalian              string = "lilalà"
	phraseJapanese             string = "リララ"
	phraseLatinAmericanSpanish string = "mimomimo"
	phraseKorean               string = "샤방"
	phraseRussian              string = "золотце"
	phraseSpanish              string = "mimomimo"
	phraseSimplifiedChinese    string = "哩啦啦"
	phraseTraditionalChinese   string = "哩啦啦"
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
