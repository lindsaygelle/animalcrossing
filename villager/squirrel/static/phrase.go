package static

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "krzzt"
	phraseCanadianFrench       string = "krzzt"
	phraseDutch                string = "bliksems"
	phraseFrench               string = "krzzt"
	phraseGerman               string = "krazz"
	phraseItalian              string = "sguscio"
	phraseJapanese             string = "ピカッ"
	phraseLatinAmericanSpanish string = "kreee"
	phraseKorean               string = "콰광"
	phraseRussian              string = "вж-ж-ж"
	phraseSpanish              string = "avería"
	phraseSimplifiedChinese    string = "闪闪"
	phraseTraditionalChinese   string = "閃閃"
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
