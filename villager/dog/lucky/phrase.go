package lucky

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "rrr-owch"
	phraseCanadianFrench       string = "graaaouch"
	phraseDutch                string = "www-⁠au"
	phraseFrench               string = "graaaouch"
	phraseGerman               string = "rrr-awau"
	phraseItalian              string = "ossorrr"
	phraseJapanese             string = "らしいよ"
	phraseLatinAmericanSpanish string = "guuuuf"
	phraseKorean               string = "그렇대"
	phraseRussian              string = "ррр-ой-ой"
	phraseSpanish              string = "guuuuAY"
	phraseSimplifiedChinese    string = "大概"
	phraseTraditionalChinese   string = "大概"
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
