package hazel

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "uni-wow"
	phraseCanadianFrench       string = "tac tac"
	phraseDutch                string = "brauw"
	phraseFrench               string = "tac tac"
	phraseGerman               string = "kekeke"
	phraseItalian              string = "blink"
	phraseJapanese             string = "だもんね"
	phraseLatinAmericanSpanish string = "tac-tac"
	phraseKorean               string = "알아알아"
	phraseRussian              string = "фундук"
	phraseSpanish              string = "castaña"
	phraseSimplifiedChinese    string = "当然哦"
	phraseTraditionalChinese   string = "當然囉"
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
