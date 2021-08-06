package victoria

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "sugar cube"
	phraseCanadianFrench       string = "susucre"
	phraseDutch                string = "klontje"
	phraseFrench               string = "susucre"
	phraseGerman               string = "gimazuka"
	phraseItalian              string = "zolletta"
	phraseJapanese             string = "いくわよ"
	phraseLatinAmericanSpanish string = "clic-cloc"
	phraseKorean               string = "달리자"
	phraseRussian              string = "сахарок"
	phraseSpanish              string = "terroncito"
	phraseSimplifiedChinese    string = "走哦"
	phraseTraditionalChinese   string = "走囉"
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
