package erik

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "chow down"
	phraseCanadianFrench       string = "breuh"
	phraseDutch                string = "schorsie"
	phraseFrench               string = "breuh"
	phraseGerman               string = "hirschl"
	phraseItalian              string = "relax"
	phraseJapanese             string = "しかじか"
	phraseLatinAmericanSpanish string = "ñume-ñume"
	phraseKorean               string = "뚜비뚜바"
	phraseRussian              string = "ем мох"
	phraseSpanish              string = "ñume-ñume"
	phraseSimplifiedChinese    string = "鹿鹿"
	phraseTraditionalChinese   string = "鹿鹿"
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
