package bones

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "yip yip"
	phraseCanadianFrench       string = "yip yip"
	phraseDutch                string = "snuf snuf"
	phraseFrench               string = "yip yip"
	phraseGerman               string = "hechel"
	phraseItalian              string = "yip yip"
	phraseJapanese             string = "ネッ"
	phraseLatinAmericanSpanish string = "yip yip"
	phraseKorean               string = "옙"
	phraseRussian              string = "тяв-тяв"
	phraseSpanish              string = "yip yip"
	phraseSimplifiedChinese    string = "对吧"
	phraseTraditionalChinese   string = "對吧"
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
