package rory

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rory"
	nameCanadianFrench       string = "Hercule"
	nameDutch                string = "Rory"
	nameFrench               string = "Hercule"
	nameGerman               string = "Leon"
	nameItalian              string = "Ruggero"
	nameJapanese             string = "アーサー"
	nameLatinAmericanSpanish string = "Rogelio"
	nameKorean               string = "아더"
	nameRussian              string = "Рори"
	nameSpanish              string = "Rogelio"
	nameSimplifiedChinese    string = "亚瑟"
	nameTraditionalChinese   string = "亞瑟"
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
