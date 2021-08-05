package chevre

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chevre"
	nameCanadianFrench       string = "Biquette"
	nameDutch                string = "Chevre"
	nameFrench               string = "Biquette"
	nameGerman               string = "Anette"
	nameItalian              string = "Diletta"
	nameJapanese             string = "ユキ"
	nameLatinAmericanSpanish string = "Cabriola"
	nameKorean               string = "윤이"
	nameRussian              string = "Шевр"
	nameSpanish              string = "Cabriola"
	nameSimplifiedChinese    string = "雪儿"
	nameTraditionalChinese   string = "雪兒"
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:      nameAmericanEnglish,
		language.CanadianFrench:       nameCanadianFrench,
		language.Dutch:                nameDutch,
		language.French:               nameFrench,
		language.German:               nameGerman,
		language.Italian:              nameItalian,
		language.Japanese:             nameJapanese,
		language.Korean:               nameKorean,
		language.LatinAmericanSpanish: nameLatinAmericanSpanish,
		language.Russian:              nameRussian,
		language.Spanish:              nameSpanish,
		language.SimplifiedChinese:    nameSimplifiedChinese,
		language.TraditionalChinese:   nameTraditionalChinese}
)
