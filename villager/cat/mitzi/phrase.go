package mitzi

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "mew"
	phraseCanadianFrench       string = "meeeeeh"
	phraseDutch                string = "mieuw"
	phraseFrench               string = "meeeeeh"
	phraseGerman               string = "muh"
	phraseItalian              string = "miao"
	phraseJapanese             string = "ニャー"
	phraseLatinAmericanSpanish string = "misi-misi"
	phraseKorean               string = "야～옹"
	phraseRussian              string = "мяу-мяу"
	phraseSpanish              string = "misi-misi"
	phraseSimplifiedChinese    string = "喵"
	phraseTraditionalChinese   string = "喵"
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
