package pompom

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "rah rah"
	phraseCanadianFrench       string = "rah rah"
	phraseDutch                string = "hup hup"
	phraseFrench               string = "rah rah"
	phraseGerman               string = "krah"
	phraseItalian              string = "quoquiquà"
	phraseJapanese             string = "だっピ"
	phraseLatinAmericanSpanish string = "ra-rá"
	phraseKorean               string = "다삐"
	phraseRussian              string = "крякушки"
	phraseSpanish              string = "sinplumas"
	phraseSimplifiedChinese    string = "裴裴"
	phraseTraditionalChinese   string = "裴裴"
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
