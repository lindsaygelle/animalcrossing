package tasha

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tasha"
	nameCanadianFrench       string = "Nadeige"
	nameDutch                string = "Tasha"
	nameFrench               string = "Nadeige"
	nameGerman               string = "Natalja"
	nameItalian              string = "Teresa"
	nameJapanese             string = "ナターシャ"
	nameLatinAmericanSpanish string = "Tania"
	nameKorean               string = "나타샤"
	nameRussian              string = "Таша"
	nameSpanish              string = "Tania"
	nameSimplifiedChinese    string = "娜塔丽"
	nameTraditionalChinese   string = "娜塔麗"
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
