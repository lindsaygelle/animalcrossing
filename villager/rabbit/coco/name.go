package coco

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Coco"
	nameCanadianFrench       string = "Coco"
	nameDutch                string = "Coco"
	nameFrench               string = "Coco"
	nameGerman               string = "Koko"
	nameItalian              string = "Coco"
	nameJapanese             string = "やよい"
	nameLatinAmericanSpanish string = "Cocoloca"
	nameKorean               string = "이요"
	nameRussian              string = "Коко"
	nameSpanish              string = "Cocoloca"
	nameSimplifiedChinese    string = "仰韶"
	nameTraditionalChinese   string = "仰韶"
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
