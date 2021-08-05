package rooney

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rooney"
	nameCanadianFrench       string = "Mike"
	nameDutch                string = "Rooney"
	nameFrench               string = "Mike"
	nameGerman               string = "Robert"
	nameItalian              string = "Balzak"
	nameJapanese             string = "マイク"
	nameLatinAmericanSpanish string = "Cerillo"
	nameKorean               string = "마이크"
	nameRussian              string = "Руни"
	nameSpanish              string = "Cerillo"
	nameSimplifiedChinese    string = "麦克"
	nameTraditionalChinese   string = "麥克"
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
