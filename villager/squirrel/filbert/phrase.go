package filbert

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "bucko"
	phraseCanadianFrench       string = "gogo"
	phraseDutch                string = "vent"
	phraseFrench               string = "gogo"
	phraseGerman               string = "manno"
	phraseItalian              string = "gnic gnic"
	phraseJapanese             string = "でしゅ"
	phraseLatinAmericanSpanish string = "guirlache"
	phraseKorean               string = "예용"
	phraseRussian              string = "дружище"
	phraseSpanish              string = "guirlache"
	phraseSimplifiedChinese    string = "是哦"
	phraseTraditionalChinese   string = "是囉"
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
