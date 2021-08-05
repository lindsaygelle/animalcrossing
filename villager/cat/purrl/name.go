package purrl

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Purrl"
	nameCanadianFrench       string = "Perle"
	nameDutch                string = "Purrl"
	nameFrench               string = "Perle"
	nameGerman               string = "Franka"
	nameItalian              string = "Felidia"
	nameJapanese             string = "たま"
	nameLatinAmericanSpanish string = "Wanda"
	nameKorean               string = "타마"
	nameRussian              string = "Перл"
	nameSpanish              string = "Wanda"
	nameSimplifiedChinese    string = "小玉"
	nameTraditionalChinese   string = "小玉"
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
