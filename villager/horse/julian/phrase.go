package julian

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "glitter"
	phraseCanadianFrench       string = "western"
	phraseDutch                string = "glitter"
	phraseFrench               string = "trottoune"
	phraseGerman               string = "wihihihi"
	phraseItalian              string = "vamos"
	phraseJapanese             string = "ね、キミ"
	phraseLatinAmericanSpanish string = "trotót"
	phraseKorean               string = "그대여"
	phraseRussian              string = "блеск"
	phraseSpanish              string = "pinchito"
	phraseSimplifiedChinese    string = "喂！你"
	phraseTraditionalChinese   string = "喂！你"
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
