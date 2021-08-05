package shep

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Shep"
	nameCanadianFrench       string = "Mehdi"
	nameDutch                string = "Shep"
	nameFrench               string = "Mehdi"
	nameGerman               string = "Thomas"
	nameItalian              string = "Frangino"
	nameJapanese             string = "ボブ"
	nameLatinAmericanSpanish string = "Fleco"
	nameKorean               string = "밥"
	nameRussian              string = "Шеп"
	nameSpanish              string = "Fleco"
	nameSimplifiedChinese    string = "包伯"
	nameTraditionalChinese   string = "包伯"
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
