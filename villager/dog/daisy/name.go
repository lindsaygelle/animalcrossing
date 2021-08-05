package daisy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Daisy"
	nameCanadianFrench       string = "Naomie"
	nameDutch                string = "Daisy"
	nameFrench               string = "Naomie"
	nameGerman               string = "Doris"
	nameItalian              string = "Fiorella"
	nameJapanese             string = "バニラ"
	nameLatinAmericanSpanish string = "Luisa"
	nameKorean               string = "바닐라"
	nameRussian              string = "Дейзи"
	nameSpanish              string = "Luisa"
	nameSimplifiedChinese    string = "香草"
	nameTraditionalChinese   string = "香草"
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
