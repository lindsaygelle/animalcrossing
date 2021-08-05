package shari

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Shari"
	nameCanadianFrench       string = "Luna"
	nameDutch                string = "Shari"
	nameFrench               string = "Luna"
	nameGerman               string = "Uta"
	nameItalian              string = "Ninì"
	nameJapanese             string = "シェリー"
	nameLatinAmericanSpanish string = "Montse"
	nameKorean               string = "젤리"
	nameRussian              string = "Шари"
	nameSpanish              string = "Montse"
	nameSimplifiedChinese    string = "雪莉"
	nameTraditionalChinese   string = "雪莉"
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
