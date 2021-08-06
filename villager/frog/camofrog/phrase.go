package camofrog

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "ten-hut"
	phraseCanadianFrench       string = "kaki"
	phraseDutch                string = "geef acht"
	phraseFrench               string = "taïaut"
	phraseGerman               string = "quak-quak"
	phraseItalian              string = "attacroak"
	phraseJapanese             string = "わい"
	phraseLatinAmericanSpanish string = "croooa"
	phraseKorean               string = "와이"
	phraseRussian              string = "шагом-квак"
	phraseSpanish              string = "recluta"
	phraseSimplifiedChinese    string = "喂"
	phraseTraditionalChinese   string = "喂"
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
