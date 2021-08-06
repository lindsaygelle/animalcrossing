package pinky

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "wah"
	phraseCanadianFrench       string = "ouah"
	phraseDutch                string = "wauw"
	phraseFrench               string = "ouah"
	phraseGerman               string = "schätzchen"
	phraseItalian              string = "cubetto"
	phraseJapanese             string = "わぉ"
	phraseLatinAmericanSpanish string = "guah"
	phraseKorean               string = "와"
	phraseRussian              string = "ну и ну"
	phraseSpanish              string = "terronillo"
	phraseSimplifiedChinese    string = "哇喔"
	phraseTraditionalChinese   string = "哇喔"
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
