package sherb

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bawwww"
	phraseCanadianFrench       string = "bââââille"
	phraseDutch                string = "oefff"
	phraseFrench               string = "bââââille"
	phraseGerman               string = "schlummer"
	phraseItalian              string = "beeewn"
	phraseJapanese             string = "ふわぁ"
	phraseLatinAmericanSpanish string = "bostezzz"
	phraseKorean               string = "흐아앙"
	phraseRussian              string = "мемеша"
	phraseSpanish              string = "bostezzz"
	phraseSimplifiedChinese    string = "轻飘"
	phraseTraditionalChinese   string = "輕飄"
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
