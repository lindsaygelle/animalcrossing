package del

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "gronk"
	phraseCanadianFrench       string = "clap"
	phraseDutch                string = "gronk"
	phraseFrench               string = "clap"
	phraseGerman               string = "happ"
	phraseItalian              string = "chomp"
	phraseJapanese             string = "プシュー"
	phraseLatinAmericanSpanish string = "chomp"
	phraseKorean               string = "처얼썩"
	phraseRussian              string = "щелк"
	phraseSpanish              string = "chomp"
	phraseSimplifiedChinese    string = "噗咻"
	phraseTraditionalChinese   string = "噗咻"
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
