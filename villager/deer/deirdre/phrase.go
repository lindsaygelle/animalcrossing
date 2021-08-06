package deirdre

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "whatevs"
	phraseCanadianFrench       string = "pan"
	phraseDutch                string = "nou ja"
	phraseFrench               string = "pan"
	phraseGerman               string = "tripptripp"
	phraseItalian              string = "babam"
	phraseJapanese             string = "まいっか"
	phraseLatinAmericanSpanish string = "uatever"
	phraseKorean               string = "됐고"
	phraseRussian              string = "как-то так"
	phraseSpanish              string = "rumiante"
	phraseSimplifiedChinese    string = "随便啦"
	phraseTraditionalChinese   string = "隨便啦"
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
