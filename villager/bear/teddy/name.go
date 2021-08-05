package teddy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Teddy"
	nameCanadianFrench       string = "Teddy"
	nameDutch                string = "Teddy"
	nameFrench               string = "Teddy"
	nameGerman               string = "Torsten"
	nameItalian              string = "Teddy"
	nameJapanese             string = "たいへいた"
	nameLatinAmericanSpanish string = "Teddy"
	nameKorean               string = "병태"
	nameRussian              string = "Тедди"
	nameSpanish              string = "Teddy"
	nameSimplifiedChinese    string = "太平"
	nameTraditionalChinese   string = "太平"
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
