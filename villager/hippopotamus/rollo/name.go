package rollo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rollo"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "que diable"
	nameGerman               string = "rülps"
	nameItalian              string = "trippo"
	nameJapanese             string = "なのな"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "hip-hip"
	nameSimplifiedChinese    string = "哈"
	nameTraditionalChinese   string = ""
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
