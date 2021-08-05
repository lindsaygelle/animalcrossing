package elmer

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Elmer"
	nameCanadianFrench       string = "Martin"
	nameDutch                string = "Elmer"
	nameFrench               string = "Martin"
	nameGerman               string = "Emil"
	nameItalian              string = "Oreste"
	nameJapanese             string = "サブレ"
	nameLatinAmericanSpanish string = "Jacinto"
	nameKorean               string = "샤브렌"
	nameRussian              string = "Элмер"
	nameSpanish              string = "Jacinto"
	nameSimplifiedChinese    string = "酥饼"
	nameTraditionalChinese   string = "酥餅"
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
