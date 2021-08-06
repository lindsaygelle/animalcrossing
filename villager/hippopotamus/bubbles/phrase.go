package bubbles

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "hipster"
	phraseCanadianFrench       string = "hip hop là"
	phraseDutch                string = "hipster"
	phraseFrench               string = "hip hop là"
	phraseGerman               string = "blupper"
	phraseItalian              string = "hiippy"
	phraseJapanese             string = "でガンス"
	phraseLatinAmericanSpanish string = "glop-glop"
	phraseKorean               string = "히포포"
	phraseRussian              string = "хипстер"
	phraseSpanish              string = "hippie"
	phraseSimplifiedChinese    string = "是是"
	phraseTraditionalChinese   string = "是是"
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
