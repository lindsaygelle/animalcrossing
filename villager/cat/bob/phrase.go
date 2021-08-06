package bob

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "pthhpth"
	phraseCanadianFrench       string = "pfioupf"
	phraseDutch                string = "poekie"
	phraseFrench               string = "pfioupf"
	phraseGerman               string = "psssps"
	phraseItalian              string = "baffo"
	phraseJapanese             string = "ネコ"
	phraseLatinAmericanSpanish string = "fishfish"
	phraseKorean               string = "고양이"
	phraseRussian              string = "кх-кх-кх"
	phraseSpanish              string = "zarpas"
	phraseSimplifiedChinese    string = "猫猫"
	phraseTraditionalChinese   string = "貓貓"
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
