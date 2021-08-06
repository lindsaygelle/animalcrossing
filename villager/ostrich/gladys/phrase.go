package gladys

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "stretch"
	phraseCanadianFrench       string = "codêt"
	phraseDutch                string = "en strek"
	phraseFrench               string = "codêt"
	phraseGerman               string = "ajajajjj"
	phraseItalian              string = "stretch"
	phraseJapanese             string = "ですの"
	phraseLatinAmericanSpanish string = "plumiplú"
	phraseKorean               string = "그렇지라"
	phraseRussian              string = "растяжка"
	phraseSpanish              string = "plumillas"
	phraseSimplifiedChinese    string = "千岁"
	phraseTraditionalChinese   string = "千歲"
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
