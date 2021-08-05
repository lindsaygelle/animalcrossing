package lopez

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lopez"
	nameCanadianFrench       string = "Jon"
	nameDutch                string = "Lopez"
	nameFrench               string = "Jon"
	nameGerman               string = "Eckart"
	nameItalian              string = "Lopez"
	nameJapanese             string = "トムソン"
	nameLatinAmericanSpanish string = "Agreste"
	nameKorean               string = "톰슨"
	nameRussian              string = "Лопес"
	nameSpanish              string = "Agreste"
	nameSimplifiedChinese    string = "汤姆"
	nameTraditionalChinese   string = "湯姆"
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
