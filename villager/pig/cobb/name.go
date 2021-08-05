package cobb

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cobb"
	nameCanadianFrench       string = "Porken"
	nameDutch                string = "Cobb"
	nameFrench               string = "Porken"
	nameGerman               string = "Rolo"
	nameItalian              string = "Wurst"
	nameJapanese             string = "ハカセ"
	nameLatinAmericanSpanish string = "Franfuá"
	nameKorean               string = "박사"
	nameRussian              string = "Кобб"
	nameSpanish              string = "Franfuá"
	nameSimplifiedChinese    string = "博士"
	nameTraditionalChinese   string = "博士"
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
