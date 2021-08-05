package elvis

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Elvis"
	nameCanadianFrench       string = "Elvis"
	nameDutch                string = "Elvis"
	nameFrench               string = "Elvis"
	nameGerman               string = "Leonardo"
	nameItalian              string = "Elvis"
	nameJapanese             string = "キング"
	nameLatinAmericanSpanish string = "Elvis"
	nameKorean               string = "킹"
	nameRussian              string = "Элвис"
	nameSpanish              string = "Elvis"
	nameSimplifiedChinese    string = "皇狮"
	nameTraditionalChinese   string = "皇獅"
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
