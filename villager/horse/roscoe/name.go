package roscoe

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Roscoe"
	nameCanadianFrench       string = "Rosco"
	nameDutch                string = "Roscoe"
	nameFrench               string = "Rosco"
	nameGerman               string = "Jolly"
	nameItalian              string = "Roscoe"
	nameJapanese             string = "シュバルツ"
	nameLatinAmericanSpanish string = "Jereza"
	nameKorean               string = "슈베르트"
	nameRussian              string = "Роско"
	nameSpanish              string = "Jereza"
	nameSimplifiedChinese    string = "黑马"
	nameTraditionalChinese   string = "黑馬"
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
