package tiara

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tiara"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "mon hippy"
	nameGerman               string = "herzchen"
	nameItalian              string = "dolcezza"
	nameJapanese             string = "てゆーか"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "amorcito"
	nameSimplifiedChinese    string = "呀嗨"
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
