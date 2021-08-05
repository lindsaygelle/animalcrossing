package boomer

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Boomer"
	nameCanadianFrench       string = "Ethan"
	nameDutch                string = "Boomer"
	nameFrench               string = "Ethan"
	nameGerman               string = "Max"
	nameItalian              string = "Icaro"
	nameJapanese             string = "ショーイ"
	nameLatinAmericanSpanish string = "Serafino"
	nameKorean               string = "팽기"
	nameRussian              string = "Бумер"
	nameSpanish              string = "Serafino"
	nameSimplifiedChinese    string = "秀益"
	nameTraditionalChinese   string = "秀益"
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
