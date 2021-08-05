package greta

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Greta"
	nameCanadianFrench       string = "Greta"
	nameDutch                string = "Greta"
	nameFrench               string = "Greta"
	nameGerman               string = "Gretel"
	nameItalian              string = "Greta"
	nameJapanese             string = "ふくこ"
	nameLatinAmericanSpanish string = "Jimena"
	nameKorean               string = "복자"
	nameRussian              string = "Грета"
	nameSpanish              string = "Jimena"
	nameSimplifiedChinese    string = "如意"
	nameTraditionalChinese   string = "如意"
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
