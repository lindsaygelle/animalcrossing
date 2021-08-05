package skye

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Skye"
	nameCanadianFrench       string = "Marilou"
	nameDutch                string = "Skye"
	nameFrench               string = "Marilou"
	nameGerman               string = "Sabine"
	nameItalian              string = "Lupilla"
	nameJapanese             string = "リリィ"
	nameLatinAmericanSpanish string = "Alderia"
	nameKorean               string = "릴리"
	nameRussian              string = "Скай"
	nameSpanish              string = "Alderia"
	nameSimplifiedChinese    string = "百合"
	nameTraditionalChinese   string = "百合"
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
