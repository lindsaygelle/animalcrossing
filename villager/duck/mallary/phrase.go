package mallary

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "quackpth"
	phraseCanadianFrench       string = "coinpff"
	phraseDutch                string = "kwaks"
	phraseFrench               string = "coinpff"
	phraseGerman               string = "quackpss"
	phraseItalian              string = "quackplà"
	phraseJapanese             string = "ヨネ"
	phraseLatinAmericanSpanish string = "recuac"
	phraseKorean               string = "모모"
	phraseRussian              string = "кряк-фи"
	phraseSpanish              string = "recuac"
	phraseSimplifiedChinese    string = "八十八"
	phraseTraditionalChinese   string = "八十八"
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
