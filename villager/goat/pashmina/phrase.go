package pashmina

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "kidders"
	phraseCanadianFrench       string = "mon moulin"
	phraseDutch                string = "bokkig"
	phraseFrench               string = "mon moulin"
	phraseGerman               string = "heckmeck"
	phraseItalian              string = "embeeeh"
	phraseJapanese             string = "かな"
	phraseLatinAmericanSpanish string = "ne-eh"
	phraseKorean               string = "그러게"
	phraseRussian              string = "козлятки"
	phraseSpanish              string = "muy bieeen"
	phraseSimplifiedChinese    string = "或许芭"
	phraseTraditionalChinese   string = "或許芭"
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
