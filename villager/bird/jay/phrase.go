package jay

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "heeeeeyy"
	phraseCanadianFrench       string = "maaaaaais"
	phraseDutch                string = "heeeeeyy"
	phraseFrench               string = "maaaaaais"
	phraseGerman               string = "hiiiijjj"
	phraseItalian              string = "cipinci"
	phraseJapanese             string = "でおます"
	phraseLatinAmericanSpanish string = "eeey"
	phraseKorean               string = "그쵸"
	phraseRussian              string = "хей-хей"
	phraseSpanish              string = "heeey"
	phraseSimplifiedChinese    string = "呢喃"
	phraseTraditionalChinese   string = "呢喃"
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
