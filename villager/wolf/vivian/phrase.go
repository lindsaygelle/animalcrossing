package vivian

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "piffle"
	phraseCanadianFrench       string = "crocroque"
	phraseDutch                string = "nonsens"
	phraseFrench               string = "crocroque"
	phraseGerman               string = "scholli"
	phraseItalian              string = "arruuu"
	phraseJapanese             string = "だわよ"
	phraseLatinAmericanSpanish string = "aú-aú"
	phraseKorean               string = "그러하다"
	phraseRussian              string = "вздор"
	phraseSpanish              string = "aú-aú"
	phraseSimplifiedChinese    string = "喔呵呵"
	phraseTraditionalChinese   string = "喔呵呵"
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
