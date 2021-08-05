package friga

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Friga"
	nameCanadianFrench       string = "Friga"
	nameDutch                string = "Friga"
	nameFrench               string = "Friga"
	nameGerman               string = "Frieda"
	nameItalian              string = "Frida"
	nameJapanese             string = "サブリナ"
	nameLatinAmericanSpanish string = "Frida"
	nameKorean               string = "사브리나"
	nameRussian              string = "Фрига"
	nameSpanish              string = "Frida"
	nameSimplifiedChinese    string = "谢宾娜"
	nameTraditionalChinese   string = "謝賓娜"
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
