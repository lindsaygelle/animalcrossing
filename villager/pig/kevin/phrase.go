package kevin

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "weeweewee"
	phraseCanadianFrench       string = "toutébon"
	phraseDutch                string = "ieieieie"
	phraseFrench               string = "toutébon"
	phraseGerman               string = "quiiiek"
	phraseItalian              string = "boink"
	phraseJapanese             string = "ウリウリ"
	phraseLatinAmericanSpanish string = "bidoink"
	phraseKorean               string = "냠냠"
	phraseRussian              string = "уи-и-и"
	phraseSpanish              string = "bidoink"
	phraseSimplifiedChinese    string = "比利比利"
	phraseTraditionalChinese   string = "比利比利"
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
