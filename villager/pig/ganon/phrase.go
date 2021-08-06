package ganon

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "heh heh"
	phraseCanadianFrench       string = "ganouïrk"
	phraseDutch                string = ""
	phraseFrench               string = "ganouïrk"
	phraseGerman               string = "ganoink"
	phraseItalian              string = "trisgrunf"
	phraseJapanese             string = "フォース"
	phraseLatinAmericanSpanish string = "ganoink"
	phraseKorean               string = "포스"
	phraseRussian              string = ""
	phraseSpanish              string = "ganoink"
	phraseSimplifiedChinese    string = ""
	phraseTraditionalChinese   string = ""
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
