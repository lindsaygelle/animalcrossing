package graham

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "indeed"
	phraseCanadianFrench       string = "au pif"
	phraseDutch                string = "inderdaad"
	phraseFrench               string = "au pif"
	phraseGerman               string = "hähäm"
	phraseItalian              string = "paella"
	phraseJapanese             string = "ですぞ"
	phraseLatinAmericanSpanish string = "ñisqui"
	phraseKorean               string = "그렇다고"
	phraseRussian              string = "резонно"
	phraseSpanish              string = "ñisqui"
	phraseSimplifiedChinese    string = "就是这样"
	phraseTraditionalChinese   string = "就是這樣"
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
